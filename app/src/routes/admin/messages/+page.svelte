<script lang="ts">
	import { onMount } from 'svelte';
	import { listResource, deleteResource, markMessageRead } from '$lib/api/admin';
	import ConfirmDialog from '$lib/components/admin/ConfirmDialog.svelte';

	let messages = $state<Record<string, unknown>[]>([]);
	let loading = $state(true);
	let selectedMessage = $state<Record<string, unknown> | null>(null);
	let confirmOpen = $state(false);
	let deleteTargetId = $state<string | null>(null);

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

	function handleDelete(id: string) {
		deleteTargetId = id;
		confirmOpen = true;
	}

	async function confirmDelete() {
		if (!deleteTargetId) return;
		confirmOpen = false;
		await deleteResource('messages', deleteTargetId);
		messages = messages.filter(m => m.id !== deleteTargetId);
		if (selectedMessage && selectedMessage.id === deleteTargetId) {
			selectedMessage = null;
		}
		deleteTargetId = null;
	}

	function cancelDelete() {
		confirmOpen = false;
		deleteTargetId = null;
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

<div class="admin-header">
	<h1>Messages</h1>
	<span class="msg-count">{messages.filter(m => !m.is_read).length} unread</span>
</div>

{#if loading}
	<div class="msg-loading">
		<div class="loading-box"></div>
		<span>Loading messages...</span>
	</div>
{:else if messages.length === 0}
	<div class="admin-card empty-state">
		<div class="empty-icon" aria-hidden="true">▪</div>
		<p class="empty-text">No Messages Yet</p>
		<p class="empty-hint">Messages from your contact form will appear here.</p>
	</div>
{:else}
	<div class="admin-card admin-table-wrap">
		<table class="admin-table">
			<thead><tr><th>Name</th><th>Email</th><th>Message</th><th>Status</th><th>Date</th><th>Actions</th></tr></thead>
			<tbody>
				{#each messages as msg}
					<tr
						class="msg-row"
						class:unread={!msg.is_read}
						onclick={() => openMessage(msg)}
					>
						<td>{msg.name}</td>
						<td class="msg-email-cell">{msg.email}</td>
						<td class="msg-preview">{String(msg.message).slice(0, 50)}{String(msg.message).length > 50 ? '...' : ''}</td>
						<td><span class={`admin-badge ${msg.is_read ? 'admin-badge-published' : 'admin-badge-draft'}`}>{msg.is_read ? 'Read' : 'Unread'}</span></td>
						<td class="msg-date">{formatDate(msg.created_at)}</td>
						<td>
							<div class="admin-actions">
								<button class="admin-btn admin-btn-sm admin-btn-danger" onclick={(e) => { e.stopPropagation(); handleDelete(String(msg.id)); }}>Delete</button>
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
	<!-- svelte-ignore a11y_interactive_supports_focus a11y_click_events_have_key_events -->
	<div class="msg-overlay" onclick={closeMessage} onkeydown={(e) => { if (e.key === 'Escape') closeMessage(); }} role="dialog" aria-modal="true" tabindex="-1">
		<div class="msg-modal" onclick={(e) => e.stopPropagation()}>
			<div class="msg-modal-stamp" aria-hidden="true">MSG</div>
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
						<a href={`mailto:${selectedMessage.email}`} class="msg-value msg-email-link">{selectedMessage.email}</a>
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
				<a href={`mailto:${selectedMessage.email}?subject=Re: Portfolio Contact`} class="admin-btn admin-btn-primary">Reply via Email</a>
				<button class="admin-btn admin-btn-danger" onclick={() => { handleDelete(String(selectedMessage?.id)); }}>Delete</button>
				<button class="admin-btn" onclick={closeMessage}>Close</button>
			</div>
		</div>
	</div>
{/if}

<ConfirmDialog
	open={confirmOpen}
	title="Hapus Pesan"
	message="Apakah kamu yakin ingin menghapus pesan ini? Tindakan ini tidak bisa dibatalkan."
	confirmText="Hapus"
	cancelText="Batal"
	variant="danger"
	onconfirm={confirmDelete}
	oncancel={cancelDelete}
/>

<style>
	.msg-count {
		font-family: var(--font-mono);
		font-size: 0.7rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
		background: var(--admin-pink);
		border: 3px solid var(--admin-ink);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
		padding: 0.3rem 0.6rem;
	}

	.msg-loading {
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

	.msg-row {
		cursor: pointer;
	}

	.msg-row.unread td {
		font-weight: 600;
		background: var(--admin-soft);
	}

	.msg-email-cell {
		font-family: var(--font-mono);
		font-size: 0.75rem;
	}

	.msg-preview {
		max-width: 200px;
		font-size: 0.82rem;
		color: var(--admin-ink-muted);
	}

	.msg-date {
		font-family: var(--font-mono);
		font-size: 0.68rem;
		white-space: nowrap;
	}

	/* Modal */
	.msg-overlay {
		position: fixed;
		inset: 0;
		background: rgba(16, 16, 16, 0.6);
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 9999;
		padding: 1rem;
		animation: overlay-in 0.15s ease;
	}

	@keyframes overlay-in {
		from { opacity: 0; }
		to { opacity: 1; }
	}

	.msg-modal {
		background: var(--admin-white);
		border: 3px solid var(--admin-ink);
		box-shadow: 8px 8px 0 0 var(--admin-ink);
		width: 100%;
		max-width: 560px;
		max-height: 80vh;
		overflow-y: auto;
		position: relative;
		animation: modal-pop 0.2s ease;
	}

	@keyframes modal-pop {
		from { transform: translateY(10px) rotate(1deg); opacity: 0; }
		to { transform: translateY(0) rotate(0); opacity: 1; }
	}

	.msg-modal-stamp {
		position: absolute;
		top: -8px;
		right: 16px;
		font-family: var(--font-mono);
		font-size: 0.58rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 1.5px;
		background: var(--admin-mint);
		border: 3px solid var(--admin-ink);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
		padding: 0.15rem 0.45rem;
	}

	.msg-modal-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 1.25rem 1.5rem;
		border-bottom: 3px solid var(--admin-ink);
	}

	.msg-modal-header h2 {
		margin: 0;
		font-family: var(--font-display);
		font-size: 1.5rem;
		letter-spacing: 1px;
		text-transform: uppercase;
	}

	.msg-close {
		background: var(--admin-pink);
		border: 3px solid var(--admin-ink);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
		width: 32px;
		height: 32px;
		display: grid;
		place-items: center;
		cursor: pointer;
		font-size: 0.9rem;
		font-weight: 700;
		transition: transform 0.1s ease;
	}

	.msg-close:hover {
		transform: translate(-2px, -2px);
		box-shadow: 5px 5px 0 0 var(--admin-ink);
	}

	.msg-modal-body {
		padding: 1.5rem;
	}

	.msg-meta {
		display: flex;
		flex-direction: column;
		gap: 0.55rem;
		margin-bottom: 1.25rem;
		padding-bottom: 1.25rem;
		border-bottom: 2px solid var(--admin-ink);
	}

	.msg-meta-row {
		display: flex;
		gap: 1rem;
		align-items: baseline;
	}

	.msg-label {
		font-family: var(--font-mono);
		font-size: 0.65rem;
		text-transform: uppercase;
		letter-spacing: 0.8px;
		color: var(--admin-ink-muted);
		min-width: 50px;
		font-weight: 700;
	}

	.msg-value {
		font-size: 0.88rem;
	}

	.msg-email-link {
		color: var(--admin-ink);
		text-decoration: none;
		border-bottom: 2px solid var(--admin-yellow);
		transition: border-color 0.1s;
	}

	.msg-email-link:hover {
		border-color: var(--admin-ink);
	}

	.msg-content {
		background: var(--admin-soft);
		border: 3px solid var(--admin-ink);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
		padding: 1.25rem;
	}

	.msg-content p {
		margin: 0;
		line-height: 1.65;
		white-space: pre-wrap;
	}

	.msg-modal-footer {
		display: flex;
		gap: 0.6rem;
		padding: 1.25rem 1.5rem;
		border-top: 3px solid var(--admin-ink);
		background: var(--admin-soft);
	}
</style>
