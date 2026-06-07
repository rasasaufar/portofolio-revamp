<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/admin/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="education" {id} title="Education" backHref="/admin/education" fields={[
	{ key: 'institution_name', label: 'Institution Name', required: true },
	{ key: 'degree', label: 'Degree', required: true },
	{ key: 'major', label: 'Major' },
	{ key: 'start_year', label: 'Start Year' },
	{ key: 'end_year', label: 'End Year' },
	{ key: 'gpa', label: 'GPA' },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'image_url', label: 'Image', type: 'image' },
	{ key: 'tags', label: 'Tags', type: 'tags', placeholder: 'Add tag...' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
