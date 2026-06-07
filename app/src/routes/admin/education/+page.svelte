<script lang="ts">
	import CrudPage from '$lib/components/admin/CrudPage.svelte';
</script>

<CrudPage
	resource="education"
	title="Education"
	fields={[
		{ key: 'institution_name', label: 'Institution', showInTable: true },
		{ key: 'degree', label: 'Degree', showInTable: true },
		{ key: 'start_year', label: 'Start Year', showInTable: true },
		{ key: 'end_year', label: 'End Year', showInTable: true }
	]}
/>
