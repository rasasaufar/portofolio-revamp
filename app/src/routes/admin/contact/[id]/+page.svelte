<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/admin/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="contact" {id} title="Contact Info" backHref="/admin/contact" fields={[
	{ key: 'email', label: 'Email' },
	{ key: 'phone', label: 'Phone' },
	{ key: 'whatsapp_url', label: 'WhatsApp URL' },
	{ key: 'github_url', label: 'GitHub URL' },
	{ key: 'linkedin_url', label: 'LinkedIn URL' },
	{ key: 'instagram_url', label: 'Instagram URL' },
	{ key: 'location', label: 'Location' },
	{ key: 'contact_description', label: 'Contact Description', type: 'textarea' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
