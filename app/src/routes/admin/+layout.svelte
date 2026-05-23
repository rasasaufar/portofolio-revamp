<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { isAuthenticated, clearToken } from '$lib/stores/auth';
	import '$lib/styles/***REMOVED***.css';

	let { children } = $props();

	const navItems = [
		{ href: '/***REMOVED***/dashboard', label: 'Dashboard', icon: '📊' },
		{ href: '/***REMOVED***/identity', label: 'Identity Console', icon: '🆔' },
		{ href: '/***REMOVED***/capabilities', label: 'Capabilities', icon: '📈' },
		{ href: '/***REMOVED***/strengths', label: 'Strengths', icon: '💪' },
		{ href: '/***REMOVED***/dossier', label: 'Dossier / About', icon: '📋' },
		{ href: '/***REMOVED***/education', label: 'Education', icon: '🎓' },
		{ href: '/***REMOVED***/experiences', label: 'Field Operations', icon: '💼' },
		{ href: '/***REMOVED***/projects', label: 'Project Lab', icon: '🔬' },
		{ href: '/***REMOVED***/certifications', label: 'Credentials', icon: '🏅' },
		{ href: '/***REMOVED***/publications', label: 'Research', icon: '📄' },
		{ href: '/***REMOVED***/contact', label: 'Contact', icon: '📞' },
		{ href: '/***REMOVED***/messages', label: 'Messages', icon: '✉️' },
		{ href: '/***REMOVED***/settings', label: 'Site Settings', icon: '⚙️' }
	];

	const currentPath = $derived($page.url.pathname);
	const isLoginPage = $derived(currentPath === '/***REMOVED***/login');

	onMount(() => {
		if (!$isAuthenticated && !isLoginPage) {
			goto('/***REMOVED***/login');
		}
	});

	function handleLogout() {
		clearToken();
		goto('/***REMOVED***/login');
	}
</script>

{#if isLoginPage}
	{@render children()}
{:else if $isAuthenticated}
	<div class="***REMOVED***-layout">
		<aside class="***REMOVED***-sidebar">
			<div class="***REMOVED***-sidebar-brand">
				<h2>// ADMIN</h2>
				<p>Portfolio CMS</p>
			</div>
			<ul class="***REMOVED***-nav">
				{#each navItems as item}
					<li>
						<a href={item.href} class:active={currentPath.startsWith(item.href)}>
							<span>{item.icon}</span>
							<span>{item.label}</span>
						</a>
					</li>
				{/each}
				<li>
					<a href="/***REMOVED***/login" onclick={(e) => { e.preventDefault(); handleLogout(); }}>
						<span>🚪</span>
						<span>Logout</span>
					</a>
				</li>
			</ul>
		</aside>
		<main class="***REMOVED***-main">
			{@render children()}
		</main>
	</div>
{:else}
	<div class="***REMOVED***-login-wrap">
		<p style="color: var(--***REMOVED***-text-muted);">Redirecting to login...</p>
	</div>
{/if}
