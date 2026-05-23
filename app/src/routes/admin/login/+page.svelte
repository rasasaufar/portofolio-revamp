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

			// Save or clear remember me
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
		<h1>// ADMIN LOGIN</h1>
		<p>Portfolio Management System</p>

		{#if error}
			<div class="***REMOVED***-error">{error}</div>
		{/if}

		<form class="***REMOVED***-form" onsubmit={handleSubmit}>
			<div class="***REMOVED***-field">
				<label for="email">Email</label>
				<input id="email" class="***REMOVED***-input" type="email" bind:value={email} required autocomplete="email" />
			</div>
			<div class="***REMOVED***-field">
				<label for="password">Password</label>
				<input id="password" class="***REMOVED***-input" type="password" bind:value={password} required autocomplete="current-password" />
			</div>
			<label style="display:flex; align-items:center; gap:0.5rem; cursor:pointer; font-size:0.85rem; color:var(--***REMOVED***-text-muted);">
				<input type="checkbox" bind:checked={rememberMe} style="accent-color:var(--***REMOVED***-accent);" />
				Ingat saya
			</label>
			<button class="***REMOVED***-btn ***REMOVED***-btn-primary" type="submit" disabled={loading} style="width:100%; justify-content:center; padding:0.75rem;">
				{loading ? 'Authenticating...' : 'Login'}
			</button>
		</form>
	</div>
</div>
