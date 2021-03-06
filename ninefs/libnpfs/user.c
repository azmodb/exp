/*
 * Copyright (C) 2005 by Latchesar Ionkov <lucho@ionkov.net>
 * Copyright (C) 2017 by Markus Sonderegger <marraison@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice (including the next
 * paragraph) shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
 * LATCHESAR IONKOV AND/OR ITS SUPPLIERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 * DEALINGS IN THE SOFTWARE.
 */
#include <sys/syscall.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <errno.h>
#include <unistd.h>
#include "npfs.h"
#include "npfsimpl.h"

void
np_user_incref(Npuser *u)
{
	if (!u)
		return;

	pthread_mutex_lock(&u->lock);
	u->refcount++;
	pthread_mutex_unlock(&u->lock);
}

void
np_user_decref(Npuser *u)
{
	int i;
	if (!u)
		return;

	pthread_mutex_lock(&u->lock);
	u->refcount--;
	if (u->refcount > 0) {
		pthread_mutex_unlock(&u->lock);
		return;
	}

	if (u->upool->udestroy)
		(*u->upool->udestroy)(u->upool, u);

	for(i = 0; i < u->ngroups; i++)
		np_group_decref(u->groups[i]);

	np_group_decref(u->dfltgroup);
	free(u->groups);

	pthread_mutex_destroy(&u->lock);
	free(u);
}

void
np_group_incref(Npgroup *g)
{
	if (!g)
		return;
	pthread_mutex_lock(&g->lock);
	g->refcount++;
	pthread_mutex_unlock(&g->lock);
}

void
np_group_decref(Npgroup *g)
{
	if (!g)
		return;

	pthread_mutex_lock(&g->lock);
	g->refcount--;
	if (g->refcount > 0) {
		pthread_mutex_unlock(&g->lock);
		return;
	}

	if (g->upool->gdestroy)
		(*g->upool->gdestroy)(g->upool, g);

	pthread_mutex_destroy(&g->lock);
	free(g);
}

static int
_setreuid(uid_t ruid, uid_t euid)
{
	return syscall(SYS_setreuid, ruid, euid);
}

static int
_setregid(uid_t rgid, uid_t egid)
{
	return syscall(SYS_setregid, rgid, egid);
}

static int
_setgroups(size_t size, gid_t *list)
{
	return syscall(SYS_setgroups, size, list);
}

int
np_change_user(Npuser *u)
{
	int i;
	gid_t *gids;

	if (_setreuid(0, 0) < 0)
		goto error;

	gids = np_malloc(u->ngroups * sizeof(gid_t));
	if (!gids)
		return -1;

	for(i = 0; i < u->ngroups; i++)
		gids[i] = u->groups[i]->gid;

	if (u->ngroups > 0) {
		if (_setgroups(u->ngroups, gids) < 0)
			goto error;
		if (_setregid(-1, u->dfltgroup->gid) < 0)
			goto error;
	}

	if (_setreuid(-1, u->uid) < 0)
		goto error;

	return 0;

error:
	np_uerror(errno);
	return -1;
}

Npuser *
np_current_user(void)
{
	return NULL;
}
