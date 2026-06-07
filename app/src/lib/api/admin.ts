/**
 * Admin API functions with JWT authentication.
 */

import { get } from 'svelte/store';
import { token, clearToken } from '$lib/stores/auth';
import { API_BASE_URL } from '$lib/config';

async function adminFetch<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
	const currentToken = get(token);
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
		...(options.headers as Record<string, string> || {})
	};

	if (currentToken) {
		headers['Authorization'] = `Bearer ${currentToken}`;
	}

	const response = await fetch(`${API_BASE_URL}${endpoint}`, {
		...options,
		headers
	});

	if (response.status === 401) {
		clearToken();
		if (typeof window !== 'undefined') {
			window.location.href = '/admin/login';
		}
		throw new Error('Session expired');
	}

	if (!response.ok) {
		const data = await response.json().catch(() => ({}));
		throw new Error(data.error || `Request failed: ${response.status}`);
	}

	if (response.status === 204) return undefined as unknown as T;
	return response.json();
}

// Auth
export async function login(email: string, password: string) {
	let res: Response;
	try {
		res = await fetch(`${API_BASE_URL}/api/auth/login`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ email, password })
		});
	} catch (err) {
		throw new Error('Cannot connect to API server. Make sure the backend is running on ' + API_BASE_URL);
	}
	if (!res.ok) {
		const data = await res.json().catch(() => ({}));
		throw new Error(data.error || `Login failed (${res.status})`);
	}
	return res.json();
}

export async function getMe() {
	return adminFetch<{ id: string; email: string; name: string }>('/api/auth/me');
}

export async function updateProfile(data: { name: string; email: string; current_password: string }) {
	return adminFetch<{ id: string; email: string; name: string }>('/api/auth/profile', {
		method: 'PUT',
		body: JSON.stringify(data)
	});
}

export async function changePassword(data: { current_password: string; new_password: string }) {
	return adminFetch<{ message: string }>('/api/auth/password', {
		method: 'PUT',
		body: JSON.stringify(data)
	});
}

// Dashboard
export async function getDashboardStats() {
	return adminFetch<Record<string, number>>('/api/admin/dashboard/stats');
}

// Generic CRUD
export async function listResource(resource: string) {
	return adminFetch<Record<string, unknown>[]>(`/api/admin/${resource}`);
}

export async function getResource(resource: string, id: string) {
	return adminFetch<Record<string, unknown>>(`/api/admin/${resource}/${id}`);
}

export async function createResource(resource: string, data: Record<string, unknown>) {
	return adminFetch<{ id: string; message: string }>(`/api/admin/${resource}`, {
		method: 'POST',
		body: JSON.stringify(data)
	});
}

export async function updateResource(resource: string, id: string, data: Record<string, unknown>) {
	return adminFetch<{ message: string }>(`/api/admin/${resource}/${id}`, {
		method: 'PUT',
		body: JSON.stringify(data)
	});
}

export async function deleteResource(resource: string, id: string) {
	return adminFetch<{ message: string }>(`/api/admin/${resource}/${id}`, {
		method: 'DELETE'
	});
}

export async function publishResource(resource: string, id: string) {
	return adminFetch<{ message: string }>(`/api/admin/${resource}/${id}/publish`, {
		method: 'PATCH'
	});
}

export async function unpublishResource(resource: string, id: string) {
	return adminFetch<{ message: string }>(`/api/admin/${resource}/${id}/unpublish`, {
		method: 'PATCH'
	});
}

export async function reorderResource(resource: string, ids: string[]) {
	return adminFetch<{ message: string }>(`/api/admin/${resource}/reorder`, {
		method: 'PUT',
		body: JSON.stringify({ ids })
	});
}

export async function markMessageRead(id: string) {
	return adminFetch<{ message: string }>(`/api/admin/messages/${id}/read`, {
		method: 'PATCH'
	});
}
