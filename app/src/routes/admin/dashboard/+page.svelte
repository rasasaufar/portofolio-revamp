<script lang="ts">
	import { onMount } from 'svelte';
	import { getDashboardStats } from '$lib/api/admin';

	let stats = $state<Record<string, number>>({});
	let loading = $state(true);

	const statLabels: Record<string, string> = {
		identity: 'Identity',
		capabilities: 'Capabilities',
		strengths: 'Strengths',
		dossier: 'Dossier',
		education: 'Education',
		experiences: 'Experiences',
		projects: 'Projects',
		certifications: 'Certifications',
		publications: 'Publications',
		contact: 'Contact',
		messages: 'Messages',
		unread_messages: 'Unread'
	};

	const statLinks: Record<string, string> = {
		identity: '/admin/identity',
		capabilities: '/admin/capabilities',
		strengths: '/admin/strengths',
		dossier: '/admin/dossier',
		education: '/admin/education',
		experiences: '/admin/experiences',
		projects: '/admin/projects',
		certifications: '/admin/certifications',
		publications: '/admin/publications',
		contact: '/admin/contact',
		messages: '/admin/messages',
		unread_messages: '/admin/messages'
	};

	onMount(async () => {
		try {
			stats = await getDashboardStats();
		} catch { /* ignore */ }
		loading = false;
	});

	function getGreeting(): string {
		const hour = new Date().getHours();
		if (hour < 12) return 'Good Morning';
		if (hour < 17) return 'Good Afternoon';
		return 'Good Evening';
	}
</script>

<div class="admin-header">
	<h1>Dashboard</h1>
	<div class="header-greeting">
		<span class="greeting-text">{getGreeting()}</span>
		<span class="greeting-date">{new Date().toLocaleDateString('id-ID', { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' })}</span>
	</div>
</div>

{#if loading}
	<div class="loading-state">
		<div class="loading-box"></div>
		<span>Loading stats...</span>
	</div>
{:else}
	<div class="admin-stats-grid">
		{#each Object.entries(stats) as [key, value], index}
			<a href={statLinks[key] ?? '#'} class="admin-stat-card" style="--i:{index}">
				<div class="stat-value">{value}</div>
				<div class="stat-label">{statLabels[key] ?? key}</div>
				<div class="stat-corner" aria-hidden="true"></div>
			</a>
		{/each}
	</div>

	<div class="dashboard-footer">
		<div class="footer-strip">
			<span>SYS.STATUS: ONLINE</span>
			<span>◈</span>
			<span>RECORDS: {Object.values(stats).reduce((a, b) => a + b, 0)}</span>
			<span>◈</span>
			<span>PORTFOLIO CMS</span>
		</div>
	</div>
{/if}

<style>
	.header-greeting {
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		gap: 0.15rem;
	}

	.greeting-text {
		font-family: var(--font-display);
		font-size: 1.3rem;
		letter-spacing: 1px;
		text-transform: uppercase;
	}

	.greeting-date {
		font-family: var(--font-mono);
		font-size: 0.65rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
		color: var(--admin-ink-muted);
	}

	.loading-state {
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

	.admin-stats-grid {
		animation: grid-enter 0.4s ease forwards;
	}

	@keyframes grid-enter {
		from { opacity: 0; transform: translateY(8px); }
		to { opacity: 1; transform: translateY(0); }
	}

	.admin-stat-card {
		text-decoration: none;
		color: inherit;
		animation: card-pop 0.3s ease forwards;
		animation-delay: calc(var(--i) * 50ms);
		opacity: 0;
	}

	@keyframes card-pop {
		from { opacity: 0; transform: translateY(6px) rotate(1deg); }
		to { opacity: 1; transform: translateY(0) rotate(0); }
	}

	.stat-corner {
		position: absolute;
		bottom: 6px;
		right: 6px;
		width: 10px;
		height: 10px;
		border: 2px solid var(--admin-ink);
		opacity: 0.4;
	}

	.dashboard-footer {
		margin-top: 2.5rem;
	}

	.footer-strip {
		border: var(--admin-border);
		box-shadow: var(--admin-shadow-sm);
		background: var(--admin-white);
		padding: 0.55rem 0.85rem;
		display: flex;
		align-items: center;
		gap: 1rem;
		font-family: var(--font-mono);
		font-size: 0.65rem;
		text-transform: uppercase;
		letter-spacing: 0.8px;
		color: var(--admin-ink-muted);
		overflow: hidden;
		white-space: nowrap;
	}
</style>
