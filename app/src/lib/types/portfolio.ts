/**
 * TypeScript interfaces matching the public API response shapes.
 */

export interface Identity {
	id: string;
	name: string;
	role: string;
	headline: string;
	description: string;
	avatar_url: string;
	current_focus: string[];
	availability_text: string;
	cta_primary_label: string;
	cta_primary_link: string;
	cta_secondary_label: string;
	cta_secondary_link: string;
}

export interface Capability {
	id: string;
	label: string;
	value: string;
	description: string;
}

export interface Strength {
	id: string;
	title: string;
	description: string;
	bullet_points: string[];
	icon_url: string;
}

export interface Dossier {
	id: string;
	title: string;
	paragraph_1: string;
	paragraph_2: string;
	paragraph_3: string;
}

export interface Education {
	id: string;
	institution_name: string;
	degree: string;
	major: string;
	start_year: string;
	end_year: string;
	gpa: string;
	description: string;
	image_url: string;
	tags: string[];
}

export interface GalleryImage {
	image: string;
	caption: string;
	description: string;
}

export interface Experience {
	id: string;
	company_name: string;
	position: string;
	start_date: string;
	end_date: string;
	is_current: boolean;
	description: string;
	bullet_points: string[];
	tech_tags: string[];
	logo_url: string;
	gallery_images: GalleryImage[];
}

export interface Project {
	id: string;
	title: string;
	category: string;
	description: string;
	tech_tags: string[];
	image_url: string;
	demo_url: string;
	repo_url: string;
	is_featured: boolean;
}

export interface Certification {
	id: string;
	title: string;
	issuer: string;
	issued_date: string;
	expired_date: string;
	credential_id: string;
	credential_url: string;
	description: string;
	skills: string[];
	image_url: string;
	category: string;
	status: string;
}

export interface Publication {
	id: string;
	title: string;
	journal_name: string;
	published_date: string;
	status: string;
	authors: string;
	description: string;
	tags: string[];
	publication_url: string;
}

export interface ContactInfo {
	id: string;
	email: string;
	phone: string;
	whatsapp_url: string;
	github_url: string;
	linkedin_url: string;
	instagram_url: string;
	location: string;
	contact_description: string;
}

export interface SiteSettings {
	site_title: string;
	meta_description: string;
	favicon_url: string;
	logo_url: string;
	footer_text: string;
	theme_mode: string;
	maintenance_mode: boolean;
}
