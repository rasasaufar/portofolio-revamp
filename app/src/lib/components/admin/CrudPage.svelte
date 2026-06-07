<script lang="ts">
	import { onMount } from 'svelte';
	import { listResource, deleteResource, publishResource, unpublishResource } from '$lib/api/admin';
	import ConfirmDialog from './ConfirmDialog.svelte';

	interface FieldDef {
		key: string;
		label: string;
		type?: 'text' | 'textarea' | 'json' | 'boolean' | 'select';
		showInTable?: boolean;
	}

	let {
		resource,
		title,
		fields,
		editHref = (id: string) => `/admin/${resource}/${id}`,
		newHref = `/admin/${resource}/new`
	}: {
		resource: string;
		title: string;
		fields: FieldDef[];
		editHref?: (id: string) => string;
		newHref?: string;
	} = $props();

	let items = $state<Record<string, unknown>[]>([]);
	let loading = $state(true);
	let toast = $state<{ type: string; message: string } | null>(null);
	let confirmOpen = $state(false);
	let deleteTargetId = $state<string | null>(null);

	const tableFields = $derived(fields.filter((f) => f.showInTable !== false));

	onMount(() => { loadData(); });

	async function loadData() {
		loading = true;
		try {
			items = await listResource(resource);
		} catch { items = []; }
		loading = false;
	}

	function handleDelete(id: string) {
		deleteTargetId = id;
		confirmOpen = true;
	}

	async function confirmDelete() {
		if (!deleteTargetId) return;
		confirmOpen = false;
		try {
			await deleteResource(resource, deleteTargetId);
			showToast('success', 'Record deleted');
			await loadData();
		} catch (e) {
			showToast('error', e instanceof Error ? e.message : 'Delete failed');
		}
		deleteTargetId = null;
	}

	function cancelDelete() {
		confirmOpen = false;
		deleteTargetId = null;
	}

	async function handlePublish(id: string, published: boolean) {
		try {
			if (published) {
				await unpublishResource(resource, id);
			} else {
				await publishResource(resource, id);
			}
			await loadData();
		} catch (e) {
			showToast('error', e instanceof Error ? e.message : 'Action failed');
		}
	}

	function showToast(type: string, message: string) {
		toast = { type, message };
		setTimeout(() => { toast = null; }, 3000);
	}

	function displayValue(item: Record<string, unknown>, key: string): string {
		const val = item[key];
		if (val === null || val === undefined) return '—';
		if (typeof val === 'boolean') return val ? '✓' : '✗';
		if (Array.isArray(val)) return val.length > 0 ? val.slice(0, 3).join(', ') + (val.length > 3 ? '...' : '') : '—';
		const str = String(val);
		return str.length > 60 ? str.slice(0, 60) + '...' : str;
	}
</script>

<div class="admin-header">
	<h1>{title}</h1>
	<a href={newHref} class="admin-btn admin-btn-primary">+ Add New</a>
</div>

{#if loading}
	<div class="crud-loading">
		<div class="loading-box"></div>
		<span>Loading records...</span>
	</div>
{:else if items.length === 0}
	<div class="admin-card empty-state">
		<div class="empty-icon" aria-hidden="true">◇</div>
		<p class="empty-text">No records found</p>
		<p class="empty-hint">Click "Add New" to create your first entry.</p>
	</div>
{:else}
	<div class="admin-card admin-table-wrap">
		<table class="admin-table">
			<thead>
				<tr>
					{#each tableFields.slice(0, 4) as field}
						<th>{field.label}</th>
					{/each}
					<th>Status</th>
					<th>Actions</th>
				</tr>
			</thead>
			<tbody>
				{#each items as item, index}
					<tr style="--row-i:{index}">
						{#each tableFields.slice(0, 4) as field}
							<td>{displayValue(item, field.key)}</td>
						{/each}
						<td>
							{#if 'is_published' in item}
								<span class={`admin-badge ${item.is_published ? 'admin-badge-published' : 'admin-badge-draft'}`}>
									{item.is_published ? 'Published' : 'Draft'}
								</span>
							{:else}
								<span class="admin-badge admin-badge-published">Active</span>
							{/if}
						</td>
						<td>
							<div class="admin-actions">
								<a href={editHref(String(item.id))} class="admin-btn admin-btn-sm">Edit</a>
								{#if 'is_published' in item}
									<button class="admin-btn admin-btn-sm" onclick={() => handlePublish(String(item.id), Boolean(item.is_published))}>
										{item.is_published ? 'Unpublish' : 'Publish'}
									</button>
								{/if}
								<button class="admin-btn admin-btn-sm admin-btn-danger" onclick={() => handleDelete(String(item.id))}>
									Delete
								</button>
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>

	<div class="record-count">
		<span>{items.length} record{items.length !== 1 ? 's' : ''}</span>
	</div>
{/if}

<ConfirmDialog
	open={confirmOpen}
	title="Hapus Data"
	message="Apakah kamu yakin ingin menghapus data ini? Tindakan ini tidak bisa dibatalkan."
	confirmText="Hapus"
	cancelText="Batal"
	variant="danger"
	onconfirm={confirmDelete}
	oncancel={cancelDelete}
/>

{#if toast}
	<div class={`admin-toast admin-toast-${toast.type}`}>{toast.message}</div>
{/if}

<style>
	.crud-loading {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		font-family: var(--font-mono);
		font-size: 0.75rem;
		text-transform: uppercase;
		color: var(--admin-ink-muted);
	}

	.loading-box {
		width: 14px;
		height: 14px;
		border: 3px solid var(--admin-ink);
		animation: spin-box 0.8s linear infinite;
	}

	@keyframes spin-box {
		from { transform: rotate(0deg); }
		to { transform: rotate(360deg); }
	}

	.empty-state {
		text-align: center;
		padding: 3rem 2rem;
	}

	.empty-icon {
		font-size: 2.5rem;
		margin-bottom: 0.75rem;
		opacity: 0.5;
	}

	.empty-text {
		font-family: var(--font-display);
		font-size: 1.5rem;
		letter-spacing: 1px;
		text-transform: uppercase;
		margin: 0;
	}

	.empty-hint {
		font-family: var(--font-mono);
		font-size: 0.72rem;
		text-transform: uppercase;
		color: var(--admin-ink-muted);
		margin: 0.5rem 0 0;
	}

	.record-count {
		margin-top: 0.75rem;
		font-family: var(--font-mono);
		font-size: 0.65rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
		color: var(--admin-ink-muted);
	}
</style>
