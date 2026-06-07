<script lang="ts">
	import CrudPage from '$lib/components/admin/CrudPage.svelte';
</script>

<CrudPage
	resource="dossier"
	title="Professional Dossier"
	fields={[
		{ key: 'title', label: 'Title', showInTable: true },
		{ key: 'paragraph_1', label: 'Paragraph 1', showInTable: true },
		{ key: 'is_published', label: 'Published', type: 'boolean', showInTable: true }
	]}
/>
