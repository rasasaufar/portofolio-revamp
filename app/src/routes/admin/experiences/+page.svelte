<script lang="ts">
	import CrudPage from '$lib/components/***REMOVED***/CrudPage.svelte';
</script>

<CrudPage
	resource="experiences"
	title="Work Experiences"
	fields={[
		{ key: 'company_name', label: 'Company', showInTable: true },
		{ key: 'position', label: 'Position', showInTable: true },
		{ key: 'start_date', label: 'Start Date', showInTable: true },
		{ key: 'end_date', label: 'End Date', showInTable: true }
	]}
/>
