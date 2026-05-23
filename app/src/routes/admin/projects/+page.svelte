<script lang="ts">
	import CrudPage from '$lib/components/***REMOVED***/CrudPage.svelte';
</script>

<CrudPage
	resource="projects"
	title="Projects"
	fields={[
		{ key: 'title', label: 'Title', showInTable: true },
		{ key: 'category', label: 'Category', showInTable: true },
		{ key: 'is_featured', label: 'Featured', type: 'boolean', showInTable: true },
		{ key: 'is_published', label: 'Published', type: 'boolean', showInTable: true }
	]}
/>
