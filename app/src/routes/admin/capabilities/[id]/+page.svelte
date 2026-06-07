<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/admin/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="capabilities" {id} title="Capabilities" backHref="/admin/capabilities" fields={[
	{ key: 'label', label: 'Label', required: true },
	{ key: 'value', label: 'Value', required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
