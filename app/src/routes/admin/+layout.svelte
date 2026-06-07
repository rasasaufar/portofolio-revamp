<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { isAuthenticated, clearToken } from '$lib/stores/auth';
	import { getDashboardStats } from '$lib/api/admin';
	import '$lib/styles/admin.css';

	let { children } = $props();
	let unreadCount = $state(0);

	const navItems = [
		{ href: '/admin/dashboard', label: 'Dashboard', icon: '◈' },
		{ href: '/admin/identity', label: 'Identity', icon: '▣' },
		{ href: '/admin/capabilities', label: 'Capabilities', icon: '△' },
		{ href: '/admin/strengths', label: 'Strengths', icon: '◆' },
		{ href: '/admin/dossier', label: 'Dossier', icon: '▤' },
		{ href: '/admin/education', label: 'Education', icon: '◎' },
		{ href: '/admin/experiences', label: 'Operations', icon: '▶' },
		{ href: '/admin/projects', label: 'Projects', icon: '⬡' },
		{ href: '/admin/certifications', label: 'Credentials', icon: '◇' },
		{ href: '/admin/publications', label: 'Research', icon: '▧' },
		{ href: '/admin/contact', label: 'Contact', icon: '◉' },
		{ href: '/admin/messages', label: 'Messages', icon: '▪' },
		{ href: '/admin/settings', label: 'Settings', icon: '⚙' }
	];

	const currentPath = $derived($page.url.pathname);
	const isLoginPage = $derived(currentPath === '/admin/login');

	onMount(() => {
		if (!$isAuthenticated && !isLoginPage) {
			goto('/admin/login');
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
		goto('/admin/login');
	}
</script>

{#if isLoginPage}
	{@render children()}
{:else if $isAuthenticated}
	<div class="admin-layout">
		<aside class="admin-sidebar">
			<div class="admin-sidebar-brand">
				<h2>Command<br/>Center</h2>
				<p>Portfolio CMS v1.0</p>
			</div>
			<ul class="admin-nav">
				{#each navItems as item}
					<li>
						<a href={item.href} class:active={currentPath.startsWith(item.href)}>
							<span class="nav-icon">{item.icon}</span>
							<span>{item.label}</span>
							{#if item.href === '/admin/messages' && unreadCount > 0}
								<span class="admin-nav-badge">{unreadCount}</span>
							{/if}
						</a>
					</li>
				{/each}
			</ul>
			<div class="admin-sidebar-bottom">
				<a href="/" class="sidebar-bottom-link" target="_blank">
					<span class="nav-icon">↗</span>
					<span>View Website</span>
				</a>
				<a href="/admin/login" class="sidebar-bottom-link logout-link" onclick={(e) => { e.preventDefault(); handleLogout(); }}>
					<span class="nav-icon">✕</span>
					<span>Logout</span>
				</a>
			</div>
		</aside>
		<main class="admin-main">
			{@render children()}
		</main>
	</div>
{:else}
	<div class="admin-login-wrap">
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
