<script lang="ts">
	import { onMount } from 'svelte';
	import { listResource, updateResource } from '$lib/api/admin';

	let formData = $state<Record<string, unknown>>({});
	let loading = $state(true);
	let saving = $state(false);
	let toast = $state('');

	onMount(async () => {
		try {
			const items = await listResource('settings');
			if (items.length > 0) formData = items[0];
		} catch {}
		loading = false;
	});

	async function handleSave(e: SubmitEvent) {
		e.preventDefault();
		saving = true;
		try {
			await updateResource('settings', String(formData.id), {
				site_title: formData.site_title,
				meta_description: formData.meta_description,
				favicon_url: formData.favicon_url,
				logo_url: formData.logo_url,
				footer_text: formData.footer_text,
				theme_mode: formData.theme_mode,
				maintenance_mode: formData.maintenance_mode
			});
			toast = 'Settings saved successfully';
			setTimeout(() => { toast = ''; }, 3000);
		} catch (e) {
			toast = 'Error: ' + (e instanceof Error ? e.message : 'Save failed');
		}
		saving = false;
	}
</script>

<div class="admin-header">
	<h1>Site Settings</h1>
</div>

{#if loading}
	<div class="settings-loading">
		<div class="loading-box"></div>
		<span>Loading settings...</span>
	</div>
{:else}
	<div class="admin-card settings-card">
		<div class="settings-stamp" aria-hidden="true">CONFIG</div>
		<form class="admin-form" onsubmit={handleSave}>
			<div class="admin-field">
				<label for="site_title">Site Title</label>
				<input id="site_title" class="admin-input" bind:value={formData.site_title as string} />
			</div>
			<div class="admin-field">
				<label for="meta_description">Meta Description</label>
				<textarea id="meta_description" class="admin-input admin-textarea" bind:value={formData.meta_description as string}></textarea>
			</div>
			<div class="admin-field">
				<label for="favicon_url">Favicon URL</label>
				<input id="favicon_url" class="admin-input" bind:value={formData.favicon_url as string} />
			</div>
			<div class="admin-field">
				<label for="logo_url">Logo URL</label>
				<input id="logo_url" class="admin-input" bind:value={formData.logo_url as string} />
			</div>
			<div class="admin-field">
				<label for="footer_text">Footer Text</label>
				<input id="footer_text" class="admin-input" bind:value={formData.footer_text as string} />
			</div>
			<div class="admin-field">
				<label for="theme_mode">Theme Mode</label>
				<select id="theme_mode" class="admin-input admin-select" bind:value={formData.theme_mode as string}>
					<option value="dark">Dark</option>
					<option value="light">Light</option>
				</select>
			</div>
			<div class="admin-field">
				<label class="maintenance-toggle">
					<input type="checkbox" bind:checked={formData.maintenance_mode as boolean} />
					<span>Maintenance Mode</span>
				</label>
			</div>
			<div class="settings-actions">
				<button class="admin-btn admin-btn-primary" type="submit" disabled={saving}>
					{saving ? '◈ Saving...' : '→ Save Settings'}
				</button>
			</div>
		</form>
	</div>
{/if}

{#if toast}
	<div class="admin-toast admin-toast-success">{toast}</div>
{/if}

<style>
	.settings-loading {
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

	.settings-card {
		position: relative;
	}

	.settings-stamp {
		position: absolute;
		top: -8px;
		right: 16px;
		font-family: var(--font-mono);
		font-size: 0.58rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 1.5px;
		background: var(--admin-yellow);
		border: 3px solid var(--admin-ink);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
		padding: 0.15rem 0.45rem;
	}

	.maintenance-toggle {
		display: flex;
		align-items: center;
		gap: 0.55rem;
		cursor: pointer;
		font-family: var(--font-mono);
		font-size: 0.75rem;
		text-transform: uppercase;
		letter-spacing: 0.3px;
	}

	.maintenance-toggle input {
		width: 18px;
		height: 18px;
		accent-color: var(--admin-ink);
		cursor: pointer;
	}

	.settings-actions {
		margin-top: 0.5rem;
		padding-top: 1rem;
		border-top: 2px solid var(--admin-ink);
	}
</style>
