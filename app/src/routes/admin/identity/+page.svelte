<script lang="ts">
	import CrudPage from '$lib/components/***REMOVED***/CrudPage.svelte';
</script>

<CrudPage
	resource="identity"
	title="Identity Console"
	fields={[
		{ key: 'name', label: 'Name', showInTable: true },
		{ key: 'role', label: 'Role', showInTable: true },
		{ key: 'headline', label: 'Headline', type: 'textarea' },
		{ key: 'availability_text', label: 'Availability', showInTable: true },
		{ key: 'is_published', label: 'Published', type: 'boolean', showInTable: true }
	]}
/>
