/**
 * Base API client for communicating with the Go backend.
 * Handles base URL configuration, JSON headers, and error handling.
 */

import { API_BASE_URL } from '$lib/config';

export class ApiError extends Error {
	status: number;
	details?: Record<string, unknown>;

	constructor(status: number, message: string, details?: Record<string, unknown>) {
		super(message);
		this.name = 'ApiError';
		this.status = status;
		this.details = details;
	}
}

interface FetchOptions {
	method?: string;
	body?: unknown;
	headers?: Record<string, string>;
}

export async function apiFetch<T>(endpoint: string, options: FetchOptions = {}): Promise<T> {
	const { method = 'GET', body, headers = {} } = options;

	const config: RequestInit = {
		method,
		headers: {
			'Content-Type': 'application/json',
			...headers
		}
	};

	if (body && method !== 'GET') {
		config.body = JSON.stringify(body);
	}

	const response = await fetch(`${API_BASE_URL}${endpoint}`, config);

	if (!response.ok) {
		let errorData: { error?: string; details?: Record<string, unknown> } = {};
		try {
			errorData = await response.json();
		} catch {
			// ignore parse errors
		}
		throw new ApiError(
			response.status,
			errorData.error || `Request failed with status ${response.status}`,
			errorData.details
		);
	}

	// Handle 204 No Content
	if (response.status === 204) {
		return undefined as unknown as T;
	}

	return response.json();
}
