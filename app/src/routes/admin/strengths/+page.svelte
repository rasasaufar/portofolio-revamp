<script lang="ts">
	import CrudPage from '$lib/components/***REMOVED***/CrudPage.svelte';
</script>

<CrudPage
	resource="strengths"
	title="Implementation Strengths"
	fields={[
		{ key: 'title', label: 'Title', showInTable: true },
		{ key: 'description', label: 'Description', showInTable: true },
		{ key: 'is_published', label: 'Published', type: 'boolean', showInTable: true }
	]}
/>
