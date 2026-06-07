<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { browser } from '$app/environment';
	import { login } from '$lib/api/admin';
	import { setToken } from '$lib/stores/auth';

	const REMEMBER_KEY = 'portfolio_admin_remember';

	let email = $state('');
	let password = $state('');
	let rememberMe = $state(false);
	let error = $state('');
	let loading = $state(false);

	onMount(() => {
		if (!browser) return;
		const saved = localStorage.getItem(REMEMBER_KEY);
		if (saved) {
			try {
				const data = JSON.parse(saved);
				email = data.email || '';
				rememberMe = true;
			} catch { /* ignore */ }
		}
	});

	async function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		error = '';
		loading = true;

		try {
			const res = await login(email, password);
			setToken(res.token);

			if (browser) {
				if (rememberMe) {
					localStorage.setItem(REMEMBER_KEY, JSON.stringify({ email }));
				} else {
					localStorage.removeItem(REMEMBER_KEY);
				}
			}

			goto('/admin/dashboard');
		} catch (err) {
			error = err instanceof Error ? err.message : 'Login failed';
		} finally {
			loading = false;
		}
	}
</script>

<div class="admin-login-wrap">
	<div class="admin-login-card">
		<div class="login-stamp" aria-hidden="true">AUTH</div>
		<h1>Admin Login</h1>
		<p>Portfolio Management System</p>

		{#if error}
			<div class="admin-error">{error}</div>
		{/if}

		<form class="admin-form" onsubmit={handleSubmit}>
			<div class="admin-field">
				<label for="email">Email</label>
				<input id="email" class="admin-input" type="email" bind:value={email} required autocomplete="email" placeholder="admin@example.com" />
			</div>
			<div class="admin-field">
				<label for="password">Password</label>
				<input id="password" class="admin-input" type="password" bind:value={password} required autocomplete="current-password" placeholder="••••••••" />
			</div>
			<label class="remember-label">
				<input type="checkbox" bind:checked={rememberMe} class="remember-check" />
				<span>Ingat saya</span>
			</label>
			<button class="admin-btn admin-btn-primary login-btn" type="submit" disabled={loading}>
				{loading ? '◈ Authenticating...' : '→ Login'}
			</button>
		</form>

		<div class="login-footer">
			<span>sys.auth // portfolio_cms</span>
		</div>
	</div>
</div>

<style>
	.login-stamp {
		position: absolute;
		top: -8px;
		right: 20px;
		font-family: var(--font-mono);
		font-size: 0.6rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 1.5px;
		background: var(--admin-mint);
		border: 3px solid var(--admin-ink);
		box-shadow: 3px 3px 0 0 var(--admin-ink);
		padding: 0.2rem 0.5rem;
	}

	.remember-label {
		display: flex;
		align-items: center;
		gap: 0.55rem;
		cursor: pointer;
		font-family: var(--font-mono);
		font-size: 0.72rem;
		text-transform: uppercase;
		letter-spacing: 0.3px;
	}

	.remember-check {
		width: 16px;
		height: 16px;
		accent-color: var(--admin-ink);
		cursor: pointer;
	}

	.login-btn {
		width: 100%;
		justify-content: center;
		padding: 0.85rem;
		font-size: 0.8rem;
		margin-top: 0.25rem;
	}

	.login-footer {
		margin-top: 1.5rem;
		padding-top: 1rem;
		border-top: 2px solid var(--admin-ink);
		font-family: var(--font-mono);
		font-size: 0.62rem;
		text-transform: uppercase;
		letter-spacing: 1px;
		color: var(--admin-ink-muted);
		text-align: center;
	}
</style>
