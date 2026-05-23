<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import '$lib/styles/neo-brutalism.css';

	let { children, data } = $props();
	const year = new Date().getFullYear();

	const isAdmin = $derived($page.url.pathname.startsWith('/***REMOVED***'));

	const siteTitle = $derived(data?.siteSettings?.site_title || 'Rasas Aufar - IT Implementation Portfolio');
	const siteDescription = $derived(data?.siteSettings?.meta_description || 'IT Implementation Professional experienced in government technology delivery, infrastructure operations, and applied data analysis.');

	const navigation = [
		{ id: 'hero', label: 'Identity' },
		{ id: 'capabilities', label: 'Capabilities' },
		{ id: 'operations', label: 'Field Operations' },
		{ id: 'laboratory', label: 'Project Lab' },
		{ id: 'credentials', label: 'Credentials' },
		{ id: 'research', label: 'Research' },
		{ id: 'contact', label: 'Contact' }
	];

	let activeSection = $state(navigation[0]?.id ?? 'hero');

	onMount(() => {
		if (isAdmin) return;

		const sections = navigation
			.map((item) => document.getElementById(item.id))
			.filter((section): section is HTMLElement => section !== null);

		if (!sections.length) return;

		const visibility = new Map<string, number>();
		for (const section of sections) {
			visibility.set(section.id, 0);
		}

		const applyHashSection = () => {
			const hashId = window.location.hash.replace('#', '');
			if (hashId && visibility.has(hashId)) {
				activeSection = hashId;
			}
		};

		applyHashSection();

		const observer = new IntersectionObserver(
			(entries) => {
				for (const entry of entries) {
					visibility.set(entry.target.id, entry.isIntersecting ? entry.intersectionRatio : 0);
				}

				let currentId = activeSection;
				let bestRatio = -1;

				for (const [id, ratio] of visibility.entries()) {
					if (ratio > bestRatio) {
						bestRatio = ratio;
						currentId = id;
					}
				}

				if (bestRatio > 0) {
					activeSection = currentId;
				}
			},
			{
				rootMargin: '-18% 0px -55% 0px',
				threshold: [0.1, 0.25, 0.45, 0.7]
			}
		);

		for (const section of sections) {
			observer.observe(section);
		}

		window.addEventListener('hashchange', applyHashSection);

		return () => {
			window.removeEventListener('hashchange', applyHashSection);
			observer.disconnect();
		};
	});
</script>

<svelte:head>
	{#if !isAdmin}
		<title>{siteTitle}</title>
		<meta name="description" content={siteDescription} />
		<link rel="icon" href="/favicon.png" />
	{/if}
</svelte:head>

{#if isAdmin}
	<!-- Admin pages render without public layout -->
	{@render children()}
{:else}
	<div class="nav-wrap">
		<div class="shell nav-content">
			<a class="brand" href="#hero" onclick={() => (activeSection = 'hero')}>
				<span class="brand-badge" aria-hidden="true"></span>
				<span class="brand-label">Rasas Aufar</span>
			</a>

			<nav class="nav-links" aria-label="Section navigation">
				{#each navigation as item}
					<a
						href={`#${item.id}`}
						class:active={activeSection === item.id}
						aria-current={activeSection === item.id ? 'page' : undefined}
						onclick={() => (activeSection = item.id)}
					>
						{item.label}
					</a>
				{/each}
			</nav>
		</div>
	</div>

	<main class="shell">
		{@render children()}

		<footer class="footer">
			<span>sys.date: {year} // rasas_aufar</span>
			<span>constructed with sveltekit</span>
		</footer>
	</main>
{/if}
