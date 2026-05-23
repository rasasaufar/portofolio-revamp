<script lang="ts">
	import { onMount } from 'svelte';
	import { listResource, deleteResource, markMessageRead } from '$lib/api/***REMOVED***';

	let messages = $state<Record<string, unknown>[]>([]);
	let loading = $state(true);

	onMount(async () => {
		try { messages = await listResource('messages'); } catch {}
		loading = false;
	});

	async function handleRead(id: string) {
		await markMessageRead(id);
		messages = messages.map(m => m.id === id ? { ...m, is_read: true } : m);
	}

	async function handleDelete(id: string) {
		if (!confirm('Delete this message?')) return;
		await deleteResource('messages', id);
		messages = messages.filter(m => m.id !== id);
	}
</script>

<div class="***REMOVED***-header">
	<h1>Contact Messages</h1>
</div>

{#if loading}
	<p style="color: var(--***REMOVED***-text-muted);">Loading...</p>
{:else if messages.length === 0}
	<div class="***REMOVED***-card"><p style="color: var(--***REMOVED***-text-muted); text-align:center; padding:2rem;">No messages yet.</p></div>
{:else}
	<div class="***REMOVED***-card ***REMOVED***-table-wrap">
		<table class="***REMOVED***-table">
			<thead><tr><th>Name</th><th>Email</th><th>Message</th><th>Status</th><th>Date</th><th>Actions</th></tr></thead>
			<tbody>
				{#each messages as msg}
					<tr style={msg.is_read ? '' : 'font-weight:600;'}>
						<td>{msg.name}</td>
						<td>{msg.email}</td>
						<td>{String(msg.message).slice(0, 50)}...</td>
						<td><span class={`***REMOVED***-badge ${msg.is_read ? '***REMOVED***-badge-published' : '***REMOVED***-badge-draft'}`}>{msg.is_read ? 'Read' : 'Unread'}</span></td>
						<td style="font-size:0.75rem;">{msg.created_at ? new Date(String(msg.created_at)).toLocaleDateString() : ''}</td>
						<td>
							<div class="***REMOVED***-actions">
								{#if !msg.is_read}<button class="***REMOVED***-btn ***REMOVED***-btn-sm" onclick={() => handleRead(String(msg.id))}>Mark Read</button>{/if}
								<button class="***REMOVED***-btn ***REMOVED***-btn-sm ***REMOVED***-btn-danger" onclick={() => handleDelete(String(msg.id))}>Delete</button>
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
{/if}
