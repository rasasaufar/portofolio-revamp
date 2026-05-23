<script lang="ts">
	import { onMount } from 'svelte';
	import { getDashboardStats } from '$lib/api/***REMOVED***';

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

	onMount(async () => {
		try {
			stats = await getDashboardStats();
		} catch { /* ignore */ }
		loading = false;
	});
</script>

<div class="***REMOVED***-header">
	<h1>Dashboard</h1>
</div>

{#if loading}
	<p style="color: var(--***REMOVED***-text-muted);">Loading stats...</p>
{:else}
	<div class="***REMOVED***-stats-grid">
		{#each Object.entries(stats) as [key, value]}
			<div class="***REMOVED***-stat-card">
				<div class="stat-value">{value}</div>
				<div class="stat-label">{statLabels[key] ?? key}</div>
			</div>
		{/each}
	</div>
{/if}
