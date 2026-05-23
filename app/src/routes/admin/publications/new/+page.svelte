<script lang="ts">
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
</script>

<FormEditor resource="publications" title="Publications" backHref="/***REMOVED***/publications" fields={[
	{ key: 'title', label: 'Title', type: 'textarea', required: true },
	{ key: 'journal_name', label: 'Journal Name', required: true },
	{ key: 'published_date', label: 'Published Date', required: true },
	{ key: 'status', label: 'Status', type: 'select', options: ['Published', 'In Review', 'Draft'] },
	{ key: 'authors', label: 'Authors', required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'tags', label: 'Tags (JSON)', type: 'json' },
	{ key: 'publication_url', label: 'Publication URL' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
