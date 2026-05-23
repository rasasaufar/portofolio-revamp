import type { LayoutServerLoad } from './$types';
import { fetchSiteSettings } from '$lib/api/public';

export const load: LayoutServerLoad = async () => {
	try {
		const siteSettings = await fetchSiteSettings();
		return { siteSettings };
	} catch {
		return {
			siteSettings: {
				site_title: 'Rasas Aufar - IT Implementation Portfolio',
				meta_description: 'IT Implementation Professional experienced in government technology delivery, infrastructure operations, and applied data analysis.',
				favicon_url: '',
				logo_url: '',
				footer_text: '',
				theme_mode: 'dark',
				maintenance_mode: false
			}
		};
	}
};
