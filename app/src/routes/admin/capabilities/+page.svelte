<script lang="ts">
	import CrudPage from '$lib/components/***REMOVED***/CrudPage.svelte';
</script>

<CrudPage
	resource="capabilities"
	title="Capabilities"
	fields={[
		{ key: 'label', label: 'Label', showInTable: true },
		{ key: 'value', label: 'Value', showInTable: true },
		{ key: 'description', label: 'Description' },
		{ key: 'is_published', label: 'Published', type: 'boolean', showInTable: true }
	]}
/>
