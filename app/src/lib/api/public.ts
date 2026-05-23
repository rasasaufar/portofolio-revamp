/**
 * Public API functions for fetching portfolio data from the backend.
 */

import { apiFetch } from './client';
import type {
	Identity,
	Capability,
	Strength,
	Dossier,
	Education,
	Experience,
	Project,
	Certification,
	Publication,
	ContactInfo,
	SiteSettings
} from '$lib/types/portfolio';

export async function fetchIdentity(): Promise<Identity | null> {
	return apiFetch<Identity | null>('/api/identity');
}

export async function fetchCapabilities(): Promise<Capability[]> {
	return apiFetch<Capability[]>('/api/capabilities');
}

export async function fetchStrengths(): Promise<Strength[]> {
	return apiFetch<Strength[]>('/api/strengths');
}

export async function fetchDossier(): Promise<Dossier | null> {
	return apiFetch<Dossier | null>('/api/dossier');
}

export async function fetchEducation(): Promise<Education[]> {
	return apiFetch<Education[]>('/api/education');
}

export async function fetchExperiences(): Promise<Experience[]> {
	return apiFetch<Experience[]>('/api/experiences');
}

export async function fetchProjects(): Promise<Project[]> {
	return apiFetch<Project[]>('/api/projects');
}

export async function fetchCertifications(): Promise<Certification[]> {
	return apiFetch<Certification[]>('/api/certifications');
}

export async function fetchPublications(): Promise<Publication[]> {
	return apiFetch<Publication[]>('/api/publications');
}

export async function fetchContactInfo(): Promise<ContactInfo | null> {
	return apiFetch<ContactInfo | null>('/api/contact-info');
}

export async function fetchSiteSettings(): Promise<SiteSettings> {
	return apiFetch<SiteSettings>('/api/site-settings');
}

export async function submitContactMessage(data: {
	name: string;
	email: string;
	message: string;
}): Promise<{ message: string }> {
	return apiFetch<{ message: string }>('/api/contact/messages', {
		method: 'POST',
		body: data
	});
}
