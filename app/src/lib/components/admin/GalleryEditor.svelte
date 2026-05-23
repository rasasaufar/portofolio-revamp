<script lang="ts">
	import ImageUpload from '$lib/components/***REMOVED***/ImageUpload.svelte';

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
				<span class="gallery-card-num">#{index + 1}</span>
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
					class="***REMOVED***-input"
					placeholder="Caption"
					value={item.caption}
					oninput={(e) => updateField(index, 'caption', (e.target as HTMLInputElement).value)}
				/>
			</div>

			<div class="gallery-field">
				<span class="gallery-label">Description</span>
				<textarea
					class="***REMOVED***-input"
					placeholder="Description"
					value={item.description}
					oninput={(e) => updateField(index, 'description', (e.target as HTMLTextAreaElement).value)}
					style="min-height:60px; resize:vertical;"
				></textarea>
			</div>
		</div>
	{/each}

	<button type="button" class="***REMOVED***-btn ***REMOVED***-btn-sm ***REMOVED***-btn-primary" onclick={addItem}>+ Add Gallery Item</button>
</div>

<style>
	.gallery-editor {
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
	}

	.gallery-card {
		background: var(--***REMOVED***-bg);
		border: 1px solid var(--***REMOVED***-border);
		border-radius: var(--***REMOVED***-radius);
		padding: 1rem;
	}

	.gallery-card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 0.75rem;
	}

	.gallery-card-num {
		font-size: 0.75rem;
		color: var(--***REMOVED***-text-muted);
		font-family: var(--***REMOVED***-mono);
	}

	.gallery-remove {
		background: none;
		border: 1px solid var(--***REMOVED***-border);
		color: var(--***REMOVED***-text-muted);
		width: 26px;
		height: 26px;
		border-radius: 4px;
		cursor: pointer;
		font-size: 0.75rem;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.gallery-remove:hover {
		background: var(--***REMOVED***-danger);
		border-color: var(--***REMOVED***-danger);
		color: white;
	}

	.gallery-field {
		margin-bottom: 0.6rem;
	}

	.gallery-label {
		display: block;
		font-size: 0.7rem;
		text-transform: uppercase;
		color: var(--***REMOVED***-text-muted);
		margin-bottom: 0.3rem;
		font-weight: 600;
	}
</style>
