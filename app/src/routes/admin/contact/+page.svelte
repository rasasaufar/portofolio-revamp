<script lang="ts">
	import CrudPage from '$lib/components/admin/CrudPage.svelte';
</script>

<CrudPage
	resource="contact"
	title="Contact Info"
	fields={[
		{ key: 'email', label: 'Email', showInTable: true },
		{ key: 'phone', label: 'Phone', showInTable: true },
		{ key: 'github_url', label: 'GitHub URL', showInTable: true },
		{ key: 'is_published', label: 'Published', type: 'boolean', showInTable: true }
	]}
/>
