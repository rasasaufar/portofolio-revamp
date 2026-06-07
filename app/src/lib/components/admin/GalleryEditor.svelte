<script lang="ts">
	import ImageUpload from '$lib/components/admin/ImageUpload.svelte';

	interface GalleryItem {
		image: string;
		caption: string;
		description: string;
	}

	let {
		value = [],
		onchange
	}: {
		value?: GalleryItem[];
		onchange?: (items: GalleryItem[]) => void;
	} = $props();

	let items = $state<GalleryItem[]>([]);

	$effect(() => {
		if (Array.isArray(value)) {
			items = value.map(v => ({ image: v.image || '', caption: v.caption || '', description: v.description || '' }));
		} else if (typeof value === 'string') {
			try { items = JSON.parse(value); } catch { items = []; }
		}
	});

	function emit() {
		onchange?.([...items]);
	}

	function addItem() {
		items = [...items, { image: '', caption: '', description: '' }];
		emit();
	}

	function removeItem(index: number) {
		items = items.filter((_, i) => i !== index);
		emit();
	}

	function updateField(index: number, field: keyof GalleryItem, val: string) {
		items[index] = { ...items[index], [field]: val };
		items = [...items];
		emit();
	}
</script>

<div class="gallery-editor">
	{#each items as item, index}
		<div class="gallery-card">
			<div class="gallery-card-header">
				<span class="gallery-card-num">Item #{index + 1}</span>
				<button type="button" class="gallery-remove" onclick={() => removeItem(index)} title="Remove">✕</button>
			</div>

			<div class="gallery-field">
				<span class="gallery-label">Image</span>
				<ImageUpload
					value={item.image}
					onchange={(url) => updateField(index, 'image', url)}
				/>
			</div>

			<div class="gallery-field">
				<span class="gallery-label">Caption</span>
				<input
					type="text"
					class="admin-input"
					placeholder="Caption"
					value={item.caption}
					oninput={(e) => updateField(index, 'caption', (e.target as HTMLInputElement).value)}
				/>
			</div>

			<div class="gallery-field">
				<span class="gallery-label">Description</span>
				<textarea
					class="admin-input"
					placeholder="Description"
					value={item.description}
					oninput={(e) => updateField(index, 'description', (e.target as HTMLTextAreaElement).value)}
					style="min-height:60px; resize:vertical;"
				></textarea>
			</div>
		</div>
	{/each}

	<button type="button" class="admin-btn admin-btn-sm admin-btn-primary" onclick={addItem}>+ Add Gallery Item</button>
</div>

<style>
	.gallery-editor {
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
	}

	.gallery-card {
		background: var(--admin-soft);
		border: 3px solid var(--admin-ink);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
		padding: 1rem;
	}

	.gallery-card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 0.75rem;
		padding-bottom: 0.5rem;
		border-bottom: 2px solid var(--admin-ink);
	}

	.gallery-card-num {
		font-family: var(--font-mono);
		font-size: 0.68rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
		color: var(--admin-ink-muted);
	}

	.gallery-remove {
		background: var(--admin-pink);
		border: 2px solid var(--admin-ink);
		box-shadow: 2px 2px 0 0 var(--admin-ink);
		width: 26px;
		height: 26px;
		cursor: pointer;
		font-size: 0.75rem;
		font-weight: 700;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: transform 0.1s ease;
	}

	.gallery-remove:hover {
		transform: translate(-1px, -1px);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
	}

	.gallery-field {
		margin-bottom: 0.6rem;
	}

	.gallery-label {
		display: block;
		font-family: var(--font-mono);
		font-size: 0.65rem;
		text-transform: uppercase;
		letter-spacing: 0.8px;
		color: var(--admin-ink-muted);
		margin-bottom: 0.3rem;
		font-weight: 700;
	}
</style>
