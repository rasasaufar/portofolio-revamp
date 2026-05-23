<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { isAuthenticated, clearToken } from '$lib/stores/auth';
	import { getDashboardStats } from '$lib/api/***REMOVED***';
	import '$lib/styles/***REMOVED***.css';

	let { children } = $props();
	let unreadCount = $state(0);

	const navItems = [
		{ href: '/***REMOVED***/dashboard', label: 'Dashboard', icon: '◈' },
		{ href: '/***REMOVED***/identity', label: 'Identity', icon: '▣' },
		{ href: '/***REMOVED***/capabilities', label: 'Capabilities', icon: '△' },
		{ href: '/***REMOVED***/strengths', label: 'Strengths', icon: '◆' },
		{ href: '/***REMOVED***/dossier', label: 'Dossier', icon: '▤' },
		{ href: '/***REMOVED***/education', label: 'Education', icon: '◎' },
		{ href: '/***REMOVED***/experiences', label: 'Operations', icon: '▶' },
		{ href: '/***REMOVED***/projects', label: 'Projects', icon: '⬡' },
		{ href: '/***REMOVED***/certifications', label: 'Credentials', icon: '◇' },
		{ href: '/***REMOVED***/publications', label: 'Research', icon: '▧' },
		{ href: '/***REMOVED***/contact', label: 'Contact', icon: '◉' },
		{ href: '/***REMOVED***/messages', label: 'Messages', icon: '▪' },
		{ href: '/***REMOVED***/settings', label: 'Settings', icon: '⚙' }
	];

	const currentPath = $derived($page.url.pathname);
	const isLoginPage = $derived(currentPath === '/***REMOVED***/login');

	onMount(() => {
		if (!$isAuthenticated && !isLoginPage) {
			goto('/***REMOVED***/login');
			return;
		}

		if ($isAuthenticated && !isLoginPage) {
			loadUnreadCount();
			const interval = setInterval(loadUnreadCount, 30000);
			return () => clearInterval(interval);
		}
	});

	async function loadUnreadCount() {
		try {
			const stats = await getDashboardStats();
			unreadCount = stats.unread_messages || 0;
		} catch { /* ignore */ }
	}

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
				<h2>Command<br/>Center</h2>
				<p>Portfolio CMS v1.0</p>
			</div>
			<ul class="***REMOVED***-nav">
				{#each navItems as item}
					<li>
						<a href={item.href} class:active={currentPath.startsWith(item.href)}>
							<span class="nav-icon">{item.icon}</span>
							<span>{item.label}</span>
							{#if item.href === '/***REMOVED***/messages' && unreadCount > 0}
								<span class="***REMOVED***-nav-badge">{unreadCount}</span>
							{/if}
						</a>
					</li>
				{/each}
				<li>
					<a href="/***REMOVED***/login" onclick={(e) => { e.preventDefault(); handleLogout(); }}>
						<span class="nav-icon">✕</span>
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
		<p style="font-family: var(--font-mono); text-transform: uppercase; font-size: 0.75rem;">Redirecting...</p>
	</div>
{/if}

<style>
	.nav-icon {
		font-size: 0.85rem;
		width: 1.2rem;
		text-align: center;
		flex-shrink: 0;
	}
</style>
