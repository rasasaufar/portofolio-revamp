<script lang="ts">
	/**
	 * Editor for simple string arrays (tags, bullet points, tech tags).
	 * Shows items as a list with add/remove buttons.
	 */

	let {
		value = [],
		onchange,
		placeholder = 'Add new item...'
	}: {
		value?: string[];
		onchange?: (items: string[]) => void;
		placeholder?: string;
	} = $props();

	let items = $state<string[]>([]);
	let newItem = $state('');

	// Sync from prop
	$effect(() => {
		if (Array.isArray(value)) {
			items = [...value];
		} else if (typeof value === 'string') {
			try { items = JSON.parse(value); } catch { items = []; }
		}
	});

	function emit() {
		onchange?.(items);
	}

	function addItem() {
		const trimmed = newItem.trim();
		if (!trimmed) return;
		items = [...items, trimmed];
		newItem = '';
		emit();
	}

	function removeItem(index: number) {
		items = items.filter((_, i) => i !== index);
		emit();
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter') {
			e.preventDefault();
			addItem();
		}
	}

	function moveUp(index: number) {
		if (index === 0) return;
		const temp = items[index];
		items[index] = items[index - 1];
		items[index - 1] = temp;
		items = [...items];
		emit();
	}

	function moveDown(index: number) {
		if (index >= items.length - 1) return;
		const temp = items[index];
		items[index] = items[index + 1];
		items[index + 1] = temp;
		items = [...items];
		emit();
	}
</script>

<div class="tag-editor">
	{#if items.length > 0}
		<ul class="tag-list">
			{#each items as item, index}
				<li class="tag-item">
					<span class="tag-index">#{index + 1}</span>
					<span class="tag-text">{item}</span>
					<div class="tag-actions">
						<button type="button" class="tag-btn" onclick={() => moveUp(index)} disabled={index === 0} title="Move up">↑</button>
						<button type="button" class="tag-btn" onclick={() => moveDown(index)} disabled={index === items.length - 1} title="Move down">↓</button>
						<button type="button" class="tag-btn tag-btn-remove" onclick={() => removeItem(index)} title="Remove">✕</button>
					</div>
				</li>
			{/each}
		</ul>
	{/if}

	<div class="tag-input-row">
		<input
			type="text"
			class="***REMOVED***-input tag-input"
			bind:value={newItem}
			onkeydown={handleKeydown}
			{placeholder}
		/>
		<button type="button" class="***REMOVED***-btn ***REMOVED***-btn-sm ***REMOVED***-btn-primary" onclick={addItem}>+ Add</button>
	</div>
</div>

<style>
	.tag-editor {
		width: 100%;
	}

	.tag-list {
		list-style: none;
		padding: 0;
		margin: 0 0 0.6rem;
		display: flex;
		flex-direction: column;
		gap: 0.35rem;
	}

	.tag-item {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		background: var(--***REMOVED***-soft);
		border: 3px solid var(--***REMOVED***-ink);
		box-shadow: 2px 2px 0 0 var(--***REMOVED***-ink);
		padding: 0.45rem 0.7rem;
		transition: transform 0.1s ease;
	}

	.tag-item:hover {
		transform: translate(-1px, -1px);
		box-shadow: 3px 3px 0 0 var(--***REMOVED***-ink);
	}

	.tag-index {
		font-family: var(--font-mono);
		font-size: 0.62rem;
		color: var(--***REMOVED***-ink-muted);
		min-width: 1.5rem;
	}

	.tag-text {
		flex: 1;
		font-size: 0.85rem;
		word-break: break-word;
	}

	.tag-actions {
		display: flex;
		gap: 0.2rem;
		flex-shrink: 0;
	}

	.tag-btn {
		background: var(--***REMOVED***-white);
		border: 2px solid var(--***REMOVED***-ink);
		color: var(--***REMOVED***-ink);
		width: 24px;
		height: 24px;
		cursor: pointer;
		font-size: 0.7rem;
		display: flex;
		align-items: center;
		justify-content: center;
		transition: all 0.1s ease;
		box-shadow: 1px 1px 0 0 var(--***REMOVED***-ink);
	}

	.tag-btn:hover:not(:disabled) {
		background: var(--***REMOVED***-yellow);
		transform: translate(-1px, -1px);
		box-shadow: 2px 2px 0 0 var(--***REMOVED***-ink);
	}

	.tag-btn:disabled {
		opacity: 0.3;
		cursor: not-allowed;
	}

	.tag-btn-remove:hover:not(:disabled) {
		background: var(--***REMOVED***-pink);
	}

	.tag-input-row {
		display: flex;
		gap: 0.5rem;
	}

	.tag-input {
		flex: 1;
	}
</style>
