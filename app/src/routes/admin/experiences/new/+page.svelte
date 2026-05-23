<script lang="ts">
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
</script>

<FormEditor resource="experiences" title="Work Experiences" backHref="/***REMOVED***/experiences" fields={[
	{ key: 'company_name', label: 'Company Name', required: true },
	{ key: 'position', label: 'Position', required: true },
	{ key: 'start_date', label: 'Start Date' },
	{ key: 'end_date', label: 'End Date' },
	{ key: 'is_current', label: 'Currently Working Here', type: 'boolean' },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'bullet_points', label: 'Bullet Points (JSON)', type: 'json' },
	{ key: 'tech_tags', label: 'Tech Tags (JSON)', type: 'json' },
	{ key: 'logo_url', label: 'Company Logo', type: 'image' },
	{ key: 'gallery_images', label: 'Gallery Images (JSON)', type: 'json' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
