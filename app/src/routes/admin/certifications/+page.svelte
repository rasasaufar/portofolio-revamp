<script lang="ts">
	import CrudPage from '$lib/components/***REMOVED***/CrudPage.svelte';
</script>

<CrudPage
	resource="certifications"
	title="Certifications"
	fields={[
		{ key: 'title', label: 'Title', showInTable: true },
		{ key: 'issuer', label: 'Issuer', showInTable: true },
		{ key: 'issued_date', label: 'Issued Date', showInTable: true },
		{ key: 'category', label: 'Category', showInTable: true }
	]}
/>
