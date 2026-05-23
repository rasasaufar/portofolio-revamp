<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="publications" {id} title="Publications" backHref="/***REMOVED***/publications" fields={[
	{ key: 'title', label: 'Title', type: 'textarea', required: true },
	{ key: 'journal_name', label: 'Journal Name', required: true },
	{ key: 'published_date', label: 'Published Date', required: true },
	{ key: 'status', label: 'Status', type: 'select', options: ['Published', 'In Review', 'Draft'] },
	{ key: 'authors', label: 'Authors', required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'tags', label: 'Tags', type: 'tags', placeholder: 'Add tag...' },
	{ key: 'publication_url', label: 'Publication URL' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
