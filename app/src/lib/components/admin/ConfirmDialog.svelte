<script lang="ts">
	let {
		open = false,
		title = 'Konfirmasi',
		message = 'Apakah kamu yakin?',
		confirmText = 'Hapus',
		cancelText = 'Batal',
		variant = 'danger',
		onconfirm,
		oncancel
	}: {
		open?: boolean;
		title?: string;
		message?: string;
		confirmText?: string;
		cancelText?: string;
		variant?: 'danger' | 'warning' | 'default';
		onconfirm?: () => void;
		oncancel?: () => void;
	} = $props();

	function handleConfirm() {
		onconfirm?.();
	}

	function handleCancel() {
		oncancel?.();
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Escape') handleCancel();
	}
</script>

{#if open}
	<!-- svelte-ignore a11y_interactive_supports_focus a11y_click_events_have_key_events -->
	<div
		class="confirm-overlay"
		onclick={handleCancel}
		onkeydown={handleKeydown}
		role="dialog"
		aria-modal="true"
		aria-labelledby="confirm-title"
		aria-describedby="confirm-message"
		tabindex="-1"
	>
		<div class="confirm-modal" onclick={(e) => e.stopPropagation()}>
			<div class="confirm-stamp" aria-hidden="true">
				{#if variant === 'danger'}⚠{:else if variant === 'warning'}⚡{:else}◆{/if}
			</div>
			<div class="confirm-header">
				<h2 id="confirm-title">{title}</h2>
			</div>
			<div class="confirm-body">
				<p id="confirm-message">{message}</p>
			</div>
			<div class="confirm-footer">
				<button class="***REMOVED***-btn" onclick={handleCancel}>{cancelText}</button>
				<button
					class="***REMOVED***-btn {variant === 'danger' ? '***REMOVED***-btn-danger' : variant === 'warning' ? '***REMOVED***-btn-warning' : '***REMOVED***-btn-primary'}"
					onclick={handleConfirm}
				>
					{confirmText}
				</button>
			</div>
		</div>
	</div>
{/if}

<style>
	.confirm-overlay {
		position: fixed;
		inset: 0;
		background: rgba(16, 16, 16, 0.6);
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 99999;
		padding: 1rem;
		animation: confirm-overlay-in 0.15s ease;
	}

	@keyframes confirm-overlay-in {
		from { opacity: 0; }
		to { opacity: 1; }
	}

	.confirm-modal {
		background: var(--***REMOVED***-white);
		border: 3px solid var(--***REMOVED***-ink);
		box-shadow: 8px 8px 0 0 var(--***REMOVED***-ink);
		width: 100%;
		max-width: 420px;
		position: relative;
		animation: confirm-pop 0.2s ease;
	}

	@keyframes confirm-pop {
		from { transform: translateY(10px) rotate(1deg); opacity: 0; }
		to { transform: translateY(0) rotate(0); opacity: 1; }
	}

	.confirm-stamp {
		position: absolute;
		top: -8px;
		right: 16px;
		font-size: 0.85rem;
		background: var(--***REMOVED***-pink);
		border: 3px solid var(--***REMOVED***-ink);
		box-shadow: 3px 3px 0 0 var(--***REMOVED***-ink);
		padding: 0.15rem 0.5rem;
		line-height: 1;
	}

	.confirm-header {
		padding: 1.25rem 1.5rem 0;
	}

	.confirm-header h2 {
		margin: 0;
		font-family: var(--font-display);
		font-size: 1.5rem;
		letter-spacing: 1px;
		text-transform: uppercase;
	}

	.confirm-body {
		padding: 1rem 1.5rem 1.5rem;
	}

	.confirm-body p {
		margin: 0;
		font-size: 0.9rem;
		line-height: 1.5;
		color: var(--***REMOVED***-ink);
	}

	.confirm-footer {
		display: flex;
		justify-content: flex-end;
		gap: 0.6rem;
		padding: 1rem 1.5rem;
		border-top: 3px solid var(--***REMOVED***-ink);
		background: var(--***REMOVED***-soft);
	}
</style>
