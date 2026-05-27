/**
 * Shared API configuration.
 * Uses the VITE_API_URL environment variable, with a local fallback.
 */
export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';
