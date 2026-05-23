<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="identity" {id} title="Identity" backHref="/***REMOVED***/identity" fields={[
	{ key: 'name', label: 'Name', required: true },
	{ key: 'role', label: 'Role', required: true },
	{ key: 'headline', label: 'Headline', type: 'textarea', required: true },
	{ key: 'description', label: 'Description', type: 'textarea', required: true },
	{ key: 'avatar_url', label: 'Avatar', type: 'image' },
	{ key: 'current_focus', label: 'Skill Stack', type: 'tags', placeholder: 'Add skill...' },
	{ key: 'availability_text', label: 'Availability Text' },
	{ key: 'cta_primary_label', label: 'CTA Primary Label' },
	{ key: 'cta_primary_link', label: 'CTA Primary Link' },
	{ key: 'cta_secondary_label', label: 'CTA Secondary Label' },
	{ key: 'cta_secondary_link', label: 'CTA Secondary Link' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
