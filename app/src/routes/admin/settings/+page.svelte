<script lang="ts">
	import { onMount } from 'svelte';
	import { changePassword, getMe, listResource, updateProfile, updateResource } from '$lib/api/admin';

	let formData = $state<Record<string, unknown>>({});
	let accountForm = $state({
		name: '',
		email: '',
		current_password: ''
	});
	let passwordForm = $state({
		current_password: '',
		new_password: '',
		confirm_password: ''
	});
	let loading = $state(true);
	let saving = $state(false);
	let savingProfile = $state(false);
	let savingPassword = $state(false);
	let toast = $state('');
	let toastKind = $state<'success' | 'error'>('success');

	onMount(async () => {
		await Promise.all([loadSiteSettings(), loadAccount()]);
		loading = false;
	});

	async function loadSiteSettings() {
		try {
			const items = await listResource('settings');
			if (items.length > 0) formData = items[0];
		} catch {
			showToast('Failed to load site settings', 'error');
		}
	}

	async function loadAccount() {
		try {
			const user = await getMe();
			accountForm.name = user.name;
			accountForm.email = user.email;
		} catch {
			showToast('Failed to load account data', 'error');
		}
	}

	function showToast(message: string, kind: 'success' | 'error' = 'success') {
		toast = message;
		toastKind = kind;
		setTimeout(() => { toast = ''; }, 3000);
	}

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
			showToast('Settings saved successfully');
		} catch (e) {
			showToast('Error: ' + (e instanceof Error ? e.message : 'Save failed'), 'error');
		}
		saving = false;
	}

	async function handleProfileSave(e: SubmitEvent) {
		e.preventDefault();
		savingProfile = true;
		try {
			const user = await updateProfile({
				name: accountForm.name,
				email: accountForm.email,
				current_password: accountForm.current_password
			});
			accountForm.name = user.name;
			accountForm.email = user.email;
			accountForm.current_password = '';
			showToast('Admin username updated successfully');
		} catch (e) {
			showToast('Error: ' + (e instanceof Error ? e.message : 'Update failed'), 'error');
		}
		savingProfile = false;
	}

	async function handlePasswordSave(e: SubmitEvent) {
		e.preventDefault();

		if (passwordForm.new_password !== passwordForm.confirm_password) {
			showToast('Error: New password confirmation does not match', 'error');
			return;
		}

		savingPassword = true;
		try {
			await changePassword({
				current_password: passwordForm.current_password,
				new_password: passwordForm.new_password
			});
			passwordForm.current_password = '';
			passwordForm.new_password = '';
			passwordForm.confirm_password = '';
			showToast('Password updated successfully');
		} catch (e) {
			showToast('Error: ' + (e instanceof Error ? e.message : 'Update failed'), 'error');
		}
		savingPassword = false;
	}
</script>

<div class="admin-header">
	<h1>Settings</h1>
</div>

{#if loading}
	<div class="settings-loading">
		<div class="loading-box"></div>
		<span>Loading settings...</span>
	</div>
{:else}
	<div class="settings-grid">
		<div class="admin-card settings-card account-card">
			<div class="settings-stamp" aria-hidden="true">ACCOUNT</div>
			<h2>Admin Account</h2>
			<form class="admin-form" onsubmit={handleProfileSave}>
				<div class="admin-field">
					<label for="admin_username">Username</label>
					<input id="admin_username" class="admin-input" bind:value={accountForm.name} required autocomplete="username" maxlength="100" />
				</div>
				<div class="admin-field">
					<label for="admin_email">Login Email</label>
					<input id="admin_email" class="admin-input" type="email" bind:value={accountForm.email} required autocomplete="email" />
				</div>
				<div class="admin-field">
					<label for="profile_current_password">Current Password</label>
					<input id="profile_current_password" class="admin-input" type="password" bind:value={accountForm.current_password} required autocomplete="current-password" />
				</div>
				<div class="settings-actions">
					<button class="admin-btn admin-btn-primary" type="submit" disabled={savingProfile}>
						{savingProfile ? '◈ Saving...' : '→ Save Account'}
					</button>
				</div>
			</form>
		</div>

		<div class="admin-card settings-card password-card">
			<div class="settings-stamp" aria-hidden="true">PASSWORD</div>
			<h2>Password</h2>
			<form class="admin-form" onsubmit={handlePasswordSave}>
				<div class="admin-field">
					<label for="password_current_password">Current Password</label>
					<input id="password_current_password" class="admin-input" type="password" bind:value={passwordForm.current_password} required autocomplete="current-password" />
				</div>
				<div class="admin-field">
					<label for="new_password">New Password</label>
					<input id="new_password" class="admin-input" type="password" bind:value={passwordForm.new_password} required minlength="6" autocomplete="new-password" />
				</div>
				<div class="admin-field">
					<label for="confirm_password">Confirm New Password</label>
					<input id="confirm_password" class="admin-input" type="password" bind:value={passwordForm.confirm_password} required minlength="6" autocomplete="new-password" />
				</div>
				<div class="settings-actions">
					<button class="admin-btn admin-btn-primary" type="submit" disabled={savingPassword}>
						{savingPassword ? '◈ Saving...' : '→ Update Password'}
					</button>
				</div>
			</form>
		</div>

		<div class="admin-card settings-card site-card">
			<div class="settings-stamp" aria-hidden="true">CONFIG</div>
			<h2>Site Settings</h2>
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
	</div>
{/if}

{#if toast}
	<div class="admin-toast" class:admin-toast-success={toastKind === 'success'} class:admin-toast-error={toastKind === 'error'}>{toast}</div>
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

	.settings-grid {
		display: grid;
		grid-template-columns: repeat(2, minmax(280px, 1fr));
		gap: 1.5rem;
		align-items: start;
	}

	.settings-card h2 {
		margin: 0 0 1.15rem;
		font-family: var(--font-display);
		font-size: 1.7rem;
		line-height: 1;
		letter-spacing: 1.5px;
		text-transform: uppercase;
	}

	.site-card {
		grid-column: 1 / -1;
	}

	.account-card {
		background: var(--admin-blue);
	}

	.password-card {
		background: var(--admin-pink);
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

	@media (max-width: 980px) {
		.settings-grid {
			grid-template-columns: 1fr;
		}

		.site-card {
			grid-column: auto;
		}
	}
</style>
