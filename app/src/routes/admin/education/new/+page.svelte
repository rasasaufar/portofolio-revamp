<script lang="ts">
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
</script>

<FormEditor resource="education" title="Education" backHref="/***REMOVED***/education" fields={[
	{ key: 'institution_name', label: 'Institution Name', required: true },
	{ key: 'degree', label: 'Degree', required: true },
	{ key: 'major', label: 'Major' },
	{ key: 'start_year', label: 'Start Year' },
	{ key: 'end_year', label: 'End Year' },
	{ key: 'gpa', label: 'GPA' },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'image_url', label: 'Image', type: 'image' },
	{ key: 'tags', label: 'Tags (JSON)', type: 'json' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
