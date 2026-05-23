<script lang="ts">
	import { onMount } from 'svelte';
	import { listResource, deleteResource, markMessageRead } from '$lib/api/***REMOVED***';

	let messages = $state<Record<string, unknown>[]>([]);
	let loading = $state(true);
	let selectedMessage = $state<Record<string, unknown> | null>(null);

	onMount(async () => {
		try { messages = await listResource('messages'); } catch {}
		loading = false;
	});

	async function openMessage(msg: Record<string, unknown>) {
		selectedMessage = msg;
		if (!msg.is_read) {
			await markMessageRead(String(msg.id));
			messages = messages.map(m => m.id === msg.id ? { ...m, is_read: true } : m);
			selectedMessage = { ...msg, is_read: true };
		}
	}

	function closeMessage() {
		selectedMessage = null;
	}

	async function handleDelete(id: string) {
		if (!confirm('Delete this message?')) return;
		await deleteResource('messages', id);
		messages = messages.filter(m => m.id !== id);
		if (selectedMessage && selectedMessage.id === id) {
			selectedMessage = null;
		}
	}

	function formatDate(dateStr: unknown): string {
		if (!dateStr) return '';
		try {
			return new Date(String(dateStr)).toLocaleString('id-ID', {
				day: '2-digit', month: 'short', year: 'numeric',
				hour: '2-digit', minute: '2-digit'
			});
		} catch { return String(dateStr); }
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
					<tr
						style={`cursor:pointer; ${msg.is_read ? '' : 'font-weight:600;'}`}
						onclick={() => openMessage(msg)}
					>
						<td>{msg.name}</td>
						<td>{msg.email}</td>
						<td>{String(msg.message).slice(0, 50)}{String(msg.message).length > 50 ? '...' : ''}</td>
						<td><span class={`***REMOVED***-badge ${msg.is_read ? '***REMOVED***-badge-published' : '***REMOVED***-badge-draft'}`}>{msg.is_read ? 'Read' : 'Unread'}</span></td>
						<td style="font-size:0.75rem;">{formatDate(msg.created_at)}</td>
						<td>
							<div class="***REMOVED***-actions">
								<button class="***REMOVED***-btn ***REMOVED***-btn-sm ***REMOVED***-btn-danger" onclick={(e) => { e.stopPropagation(); handleDelete(String(msg.id)); }}>Delete</button>
							</div>
						</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
{/if}

<!-- Message Detail Modal -->
{#if selectedMessage}
	<div class="msg-overlay" onclick={closeMessage} role="dialog" aria-modal="true">
		<div class="msg-modal" onclick={(e) => e.stopPropagation()}>
			<div class="msg-modal-header">
				<h2>Message Detail</h2>
				<button class="msg-close" onclick={closeMessage} aria-label="Close">✕</button>
			</div>
			<div class="msg-modal-body">
				<div class="msg-meta">
					<div class="msg-meta-row">
						<span class="msg-label">From</span>
						<span class="msg-value">{selectedMessage.name}</span>
					</div>
					<div class="msg-meta-row">
						<span class="msg-label">Email</span>
						<a href={`mailto:${selectedMessage.email}`} class="msg-value msg-email">{selectedMessage.email}</a>
					</div>
					<div class="msg-meta-row">
						<span class="msg-label">Date</span>
						<span class="msg-value">{formatDate(selectedMessage.created_at)}</span>
					</div>
				</div>
				<div class="msg-content">
					<p>{selectedMessage.message}</p>
				</div>
			</div>
			<div class="msg-modal-footer">
				<a href={`mailto:${selectedMessage.email}?subject=Re: Portfolio Contact`} class="***REMOVED***-btn ***REMOVED***-btn-primary">Reply via Email</a>
				<button class="***REMOVED***-btn ***REMOVED***-btn-danger" onclick={() => { handleDelete(String(selectedMessage?.id)); }}>Delete</button>
				<button class="***REMOVED***-btn" onclick={closeMessage}>Close</button>
			</div>
		</div>
	</div>
{/if}

<style>
	.msg-overlay {
		position: fixed;
		inset: 0;
		background: rgba(0, 0, 0, 0.7);
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 9999;
		padding: 1rem;
	}

	.msg-modal {
		background: var(--***REMOVED***-surface);
		border: 1px solid var(--***REMOVED***-border);
		border-radius: var(--***REMOVED***-radius);
		width: 100%;
		max-width: 550px;
		max-height: 80vh;
		overflow-y: auto;
	}

	.msg-modal-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1.25rem 1.5rem;
		border-bottom: 1px solid var(--***REMOVED***-border);
	}

	.msg-modal-header h2 {
		margin: 0;
		font-size: 1.1rem;
		font-family: var(--***REMOVED***-mono);
	}

	.msg-close {
		background: none;
		border: none;
		color: var(--***REMOVED***-text-muted);
		font-size: 1.2rem;
		cursor: pointer;
		padding: 0.25rem;
	}

	.msg-close:hover {
		color: var(--***REMOVED***-text);
	}

	.msg-modal-body {
		padding: 1.5rem;
	}

	.msg-meta {
		display: flex;
		flex-direction: column;
		gap: 0.6rem;
		margin-bottom: 1.25rem;
		padding-bottom: 1.25rem;
		border-bottom: 1px solid var(--***REMOVED***-border);
	}

	.msg-meta-row {
		display: flex;
		gap: 1rem;
		align-items: baseline;
	}

	.msg-label {
		font-size: 0.75rem;
		text-transform: uppercase;
		color: var(--***REMOVED***-text-muted);
		min-width: 50px;
		font-weight: 600;
	}

	.msg-value {
		color: var(--***REMOVED***-text);
		font-size: 0.9rem;
	}

	.msg-email {
		color: var(--***REMOVED***-accent);
		text-decoration: none;
	}

	.msg-email:hover {
		text-decoration: underline;
	}

	.msg-content {
		background: var(--***REMOVED***-bg);
		border: 1px solid var(--***REMOVED***-border);
		border-radius: var(--***REMOVED***-radius);
		padding: 1.25rem;
	}

	.msg-content p {
		margin: 0;
		line-height: 1.6;
		white-space: pre-wrap;
		color: var(--***REMOVED***-text);
	}

	.msg-modal-footer {
		display: flex;
		gap: 0.6rem;
		padding: 1.25rem 1.5rem;
		border-top: 1px solid var(--***REMOVED***-border);
	}
</style>
