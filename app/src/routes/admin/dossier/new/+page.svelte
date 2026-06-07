<script lang="ts">
	import FormEditor from '$lib/components/admin/FormEditor.svelte';
</script>

<FormEditor resource="dossier" title="Professional Dossier" backHref="/admin/dossier" fields={[
	{ key: 'title', label: 'Title', required: true },
	{ key: 'paragraph_1', label: 'Paragraph 1', type: 'textarea', required: true },
	{ key: 'paragraph_2', label: 'Paragraph 2', type: 'textarea' },
	{ key: 'paragraph_3', label: 'Paragraph 3', type: 'textarea' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
