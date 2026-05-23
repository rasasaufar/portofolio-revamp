<script lang="ts">
	import { page } from '$app/stores';
	import FormEditor from '$lib/components/***REMOVED***/FormEditor.svelte';
	const id = $derived($page.params.id);
</script>

<FormEditor resource="strengths" {id} title="Implementation Strengths" backHref="/***REMOVED***/strengths" fields={[
	{ key: 'title', label: 'Title', required: true },
	{ key: 'description', label: 'Description', type: 'textarea' },
	{ key: 'bullet_points', label: 'Bullet Points (JSON)', type: 'json' },
	{ key: 'icon_url', label: 'Icon', type: 'image' },
	{ key: 'order_number', label: 'Order', type: 'number' },
	{ key: 'is_published', label: 'Published', type: 'boolean' }
]} />
