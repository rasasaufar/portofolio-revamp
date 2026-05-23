<script lang="ts">
	import { onMount } from 'svelte';
	import { listResource, deleteResource, publishResource, unpublishResource } from '$lib/api/***REMOVED***';

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
		editHref = (id: string) => `/***REMOVED***/${resource}/${id}`,
		newHref = `/***REMOVED***/${resource}/new`
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

	const tableFields = $derived(fields.filter((f) => f.showInTable !== false));

	onMount(() => { loadData(); });

	async function loadData() {
		loading = true;
		try {
			items = await listResource(resource);
		} catch { items = []; }
		loading = false;
	}

	async function handleDelete(id: string) {
		if (!confirm('Are you sure you want to delete this record?')) return;
		try {
			await deleteResource(resource, id);
			showToast('success', 'Record deleted');
			await loadData();
		} catch (e) {
			showToast('error', e instanceof Error ? e.message : 'Delete failed');
		}
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

<div class="***REMOVED***-header">
	<h1>{title}</h1>
	<a href={newHref} class="***REMOVED***-btn ***REMOVED***-btn-primary">+ Add New</a>
</div>

{#if loading}
	<p style="color: var(--***REMOVED***-text-muted);">Loading...</p>
{:else if items.length === 0}
	<div class="***REMOVED***-card">
		<p style="color: var(--***REMOVED***-text-muted); text-align: center; padding: 2rem;">No records found. Click "Add New" to create one.</p>
	</div>
{:else}
	<div class="***REMOVED***-card ***REMOVED***-table-wrap">
		<table class="***REMOVED***-table">
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
				{#each items as item}
					<tr>
						{#each tableFields.slice(0, 4) as field}
							<td>{displayValue(item, field.key)}</td>
						{/each}
						<td>
							{#if 'is_published' in item}
								<span class={`***REMOVED***-badge ${item.is_published ? '***REMOVED***-badge-published' : '***REMOVED***-badge-draft'}`}>
									{item.is_published ? 'Published' : 'Draft'}
								</span>
							{:else}
								<span class="***REMOVED***-badge ***REMOVED***-badge-published">Active</span>
							{/if}
						</td>
						<td>
							<div class="***REMOVED***-actions">
								<a href={editHref(String(item.id))} class="***REMOVED***-btn ***REMOVED***-btn-sm">Edit</a>
								{#if 'is_published' in item}
									<button class="***REMOVED***-btn ***REMOVED***-btn-sm" onclick={() => handlePublish(String(item.id), Boolean(item.is_published))}>
										{item.is_published ? 'Unpublish' : 'Publish'}
									</button>
								{/if}
								<button class="***REMOVED***-btn ***REMOVED***-btn-sm ***REMOVED***-btn-danger" onclick={() => handleDelete(String(item.id))}>
									Delete
								</button>
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
{/if}

{#if toast}
	<div class={`***REMOVED***-toast ***REMOVED***-toast-${toast.type}`}>{toast.message}</div>
{/if}
