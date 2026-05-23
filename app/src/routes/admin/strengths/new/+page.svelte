<script lang="ts">
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
</script>

<FormEditor resource="strengths" title="Implementation Strengths" backHref="/***REMOVED***/strengths" fields={[
	{ key: 'title', label: 'Title', required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'bullet_points', label: 'Bullet Points (JSON)', type: 'json' },
	{ key: 'icon_url', label: 'Icon', type: 'image' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
