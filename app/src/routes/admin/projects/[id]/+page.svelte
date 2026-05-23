<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="projects" {id} title="Projects" backHref="/***REMOVED***/projects" fields={[
	{ key: 'title', label: 'Title', required: true },
	{ key: 'category', label: 'Category', type: 'select', options: ['Infrastructure', 'UI/UX'], required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'tech_tags', label: 'Tech Tags (JSON)', type: 'json' },
	{ key: 'image_url', label: 'Project Image', type: 'image' },
	{ key: 'demo_url', label: 'Demo URL' },
	{ key: 'repo_url', label: 'Repo URL' },
	{ key: 'is_featured', label: 'Featured', type: 'boolean' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
