<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { browser } from '$app/environment';
	import { login } from '$lib/api/***REMOVED***';
	import { setToken } from '$lib/stores/auth';

	const REMEMBER_KEY = 'portfolio_***REMOVED***_remember';

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

			goto('/***REMOVED***/dashboard');
		} catch (err) {
			error = err instanceof Error ? err.message : 'Login failed';
		} finally {
			loading = false;
		}
	}
</script>

<div class="***REMOVED***-login-wrap">
	<div class="***REMOVED***-login-card">
		<div class="login-stamp" aria-hidden="true">AUTH</div>
		<h1>Admin Login</h1>
		<p>Portfolio Management System</p>

		{#if error}
			<div class="***REMOVED***-error">{error}</div>
		{/if}

		<form class="***REMOVED***-form" onsubmit={handleSubmit}>
			<div class="***REMOVED***-field">
				<label for="email">Email</label>
				<input id="email" class="***REMOVED***-input" type="email" bind:value={email} required autocomplete="email" placeholder="***REMOVED***@example.com" />
			</div>
			<div class="***REMOVED***-field">
				<label for="password">Password</label>
				<input id="password" class="***REMOVED***-input" type="password" bind:value={password} required autocomplete="current-password" placeholder="••••••••" />
			</div>
			<label class="remember-label">
				<input type="checkbox" bind:checked={rememberMe} class="remember-check" />
				<span>Ingat saya</span>
			</label>
			<button class="***REMOVED***-btn ***REMOVED***-btn-primary login-btn" type="submit" disabled={loading}>
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
		background: var(--***REMOVED***-mint);
		border: 3px solid var(--***REMOVED***-ink);
		box-shadow: 3px 3px 0 0 var(--***REMOVED***-ink);
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
		accent-color: var(--***REMOVED***-ink);
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
		border-top: 2px solid var(--***REMOVED***-ink);
		font-family: var(--font-mono);
		font-size: 0.62rem;
		text-transform: uppercase;
		letter-spacing: 1px;
		color: var(--***REMOVED***-ink-muted);
		text-align: center;
	}
</style>
