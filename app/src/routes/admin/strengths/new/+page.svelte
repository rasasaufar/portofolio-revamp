<script lang="ts">
	import FormEditor from '$lib/components/admin/FormEditor.svelte';
</script>

<FormEditor resource="strengths" title="Implementation Strengths" backHref="/admin/strengths" fields={[
	{ key: 'title', label: 'Title', required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'bullet_points', label: 'Bullet Points', type: 'tags', placeholder: 'Add point...' },
	{ key: 'icon_url', label: 'Icon', type: 'image' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
