import { writable, derived } from 'svelte/store';
import { browser } from '$app/environment';

const TOKEN_KEY = 'portfolio_***REMOVED***_token';

function getStoredToken(): string | null {
	if (!browser) return null;
	return localStorage.getItem(TOKEN_KEY);
}

export const token = writable<string | null>(getStoredToken());

export const isAuthenticated = derived(token, ($token) => !!$token);

export function setToken(newToken: string) {
	if (browser) localStorage.setItem(TOKEN_KEY, newToken);
	token.set(newToken);
}

export function clearToken() {
	if (browser) localStorage.removeItem(TOKEN_KEY);
	token.set(null);
}
