import type { PageServerLoad } from './$types';
import {
	fetchIdentity,
	fetchCapabilities,
	fetchStrengths,
	fetchDossier,
	fetchEducation,
	fetchExperiences,
	fetchProjects,
	fetchCertifications,
	fetchPublications,
	fetchContactInfo,
	fetchSiteSettings
} from '$lib/api/public';
import { portfolioData } from '$lib/data/portfolio';

export const load: PageServerLoad = async () => {
	try {
		const [
			identity,
			capabilities,
			strengths,
			dossier,
			education,
			experiences,
			projects,
			certifications,
			publications,
			contactInfo,
			siteSettings
		] = await Promise.all([
			fetchIdentity().catch(() => null),
			fetchCapabilities().catch(() => []),
			fetchStrengths().catch(() => []),
			fetchDossier().catch(() => null),
			fetchEducation().catch(() => []),
			fetchExperiences().catch(() => []),
			fetchProjects().catch(() => []),
			fetchCertifications().catch(() => []),
			fetchPublications().catch(() => []),
			fetchContactInfo().catch(() => null),
			fetchSiteSettings().catch(() => null)
		]);

		// Check if API returned valid data (at least identity exists)
		const apiAvailable = identity !== null && identity !== undefined;

		if (apiAvailable) {
			return {
				source: 'api' as const,
				identity,
				capabilities,
				strengths,
				dossier,
				education,
				experiences,
				projects,
				certifications,
				publications,
				contactInfo,
				siteSettings
			};
		}

		// Fallback to static data
		return buildStaticFallback();
	} catch {
		// Full fallback to static data if API is completely unreachable
		return buildStaticFallback();
	}
};

function buildStaticFallback() {
	return {
		source: 'static' as const,
		identity: {
			id: '',
			name: portfolioData.profile.name,
			role: portfolioData.profile.role,
			headline: portfolioData.profile.lead,
			description: portfolioData.profile.lead,
			avatar_url: '',
			current_focus: portfolioData.skillStack,
			availability_text: portfolioData.profile.availability,
			cta_primary_label: 'View Field Operations',
			cta_primary_link: '#operations',
			cta_secondary_label: 'Open Contact Channel',
			cta_secondary_link: '#contact'
		},
		capabilities: portfolioData.stats.map((s, i) => ({
			id: String(i),
			label: s.label,
			value: s.value,
			description: ''
		})),
		strengths: portfolioData.capabilityAreas.map((a, i) => ({
			id: String(i),
			title: a.title,
			description: '',
			bullet_points: a.points,
			icon_url: ''
		})),
		dossier: {
			id: '',
			title: 'Professional Dossier',
			paragraph_1: portfolioData.profile.biography[0] || '',
			paragraph_2: portfolioData.profile.biography[1] || '',
			paragraph_3: portfolioData.profile.biography[2] || ''
		},
		education: portfolioData.education.map((e, i) => ({
			id: String(i),
			institution_name: e.institution,
			degree: e.degree,
			major: '',
			start_year: e.period.split(' - ')[0] || '',
			end_year: e.period.split(' - ')[1] || '',
			gpa: e.gpa || '',
			description: e.description,
			image_url: e.logo,
			tags: e.highlights
		})),
		experiences: portfolioData.experiences.map((e, i) => ({
			id: String(i),
			company_name: e.company,
			position: e.position,
			start_date: e.period.split(' - ')[0] || '',
			end_date: e.period.split(' - ')[1] || '',
			is_current: e.period.includes('Present'),
			description: e.description,
			bullet_points: e.responsibilities,
			tech_tags: e.tech,
			logo_url: e.logo || '',
			gallery_images: e.gallery.map((g) => ({
				image: g.image || '',
				caption: g.caption,
				description: g.description
			}))
		})),
		projects: portfolioData.projects.map((p, i) => ({
			id: String(i),
			title: p.name,
			category: p.category,
			description: p.description,
			tech_tags: p.tech,
			image_url: p.image || '',
			demo_url: p.demo || '',
			repo_url: p.github || '',
			is_featured: i < 5
		})),
		certifications: portfolioData.certificates.map((c, i) => ({
			id: String(i),
			title: c.name,
			issuer: c.issuer,
			issued_date: c.year,
			expired_date: '',
			credential_id: c.id || '',
			credential_url: '',
			description: c.description,
			skills: c.skills,
			image_url: c.image,
			category: c.category,
			status: c.verified ? 'verified' : 'completed'
		})),
		publications: portfolioData.publications.map((p, i) => ({
			id: String(i),
			title: p.title,
			journal_name: p.journal,
			published_date: p.year,
			status: p.status,
			authors: p.authors,
			description: p.abstract,
			tags: p.tags,
			publication_url: p.doi
		})),
		contactInfo: {
			id: '',
			email: 'rasasaufar4@gmail.com',
			phone: '(+62) 85326775595',
			whatsapp_url: 'https://wa.me/085326775595',
			github_url: 'https://github.com/rasasaufar',
			linkedin_url: 'https://www.linkedin.com/in/rasasaufar/',
			instagram_url: 'https://instagram.com/rasasaufar',
			location: 'Indonesia',
			contact_description: ''
		},
		siteSettings: {
			site_title: portfolioData.meta.title,
			meta_description: portfolioData.meta.description,
			favicon_url: '',
			logo_url: '',
			footer_text: '',
			theme_mode: 'dark',
			maintenance_mode: false
		}
	};
}
