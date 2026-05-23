<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="certifications" {id} title="Certifications" backHref="/***REMOVED***/certifications" fields={[
	{ key: 'title', label: 'Title', required: true },
	{ key: 'issuer', label: 'Issuer', required: true },
	{ key: 'issued_date', label: 'Issued Date', required: true },
	{ key: 'expired_date', label: 'Expired Date' },
	{ key: 'credential_id', label: 'Credential ID' },
	{ key: 'credential_url', label: 'Credential URL' },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'skills', label: 'Skills', type: 'tags', placeholder: 'Add skill...' },
	{ key: 'image_url', label: 'Certificate Image', type: 'image' },
	{ key: 'category', label: 'Category', type: 'select', options: ['core', 'additional'] },
	{ key: 'status', label: 'Status', type: 'select', options: ['active', 'expired', 'revoked'] },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
