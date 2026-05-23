<script lang="ts">
	import CrudPage from '$lib/components/***REMOVED***/CrudPage.svelte';
</script>

<CrudPage
	resource="publications"
	title="Publications"
	fields={[
		{ key: 'title', label: 'Title', showInTable: true },
		{ key: 'journal_name', label: 'Journal', showInTable: true },
		{ key: 'published_date', label: 'Published Date', showInTable: true },
		{ key: 'status', label: 'Status', showInTable: true }
	]}
/>
