<script lang="ts">
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
</script>

<FormEditor resource="capabilities" title="Capabilities" backHref="/***REMOVED***/capabilities" fields={[
	{ key: 'label', label: 'Label', required: true },
	{ key: 'value', label: 'Value', required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
