<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import '$lib/styles/neo-brutalism.css';

	let { children, data } = $props();
	const year = new Date().getFullYear();

	const isAdmin = $derived($page.url.pathname.startsWith('/admin'));
	const isMaintenanceMode = $derived(!isAdmin && data?.siteSettings?.maintenance_mode === true);

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
	let navScrolled = $state(false);
	let showBackToTop = $state(false);

	function scrollToTop() {
		window.scrollTo({ top: 0, behavior: 'smooth' });
	}

	onMount(() => {
		if (isAdmin) return;

		// Navbar scroll effect
		const handleScroll = () => {
			navScrolled = window.scrollY > 40;
			showBackToTop = window.scrollY > 600;
		};

		handleScroll(); // initial check
		window.addEventListener('scroll', handleScroll, { passive: true });

		const sections = navigation
			.map((item) => document.getElementById(item.id))
			.filter((section): section is HTMLElement => section !== null);

		if (!sections.length) {
			return () => { window.removeEventListener('scroll', handleScroll); };
		}

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
			window.removeEventListener('scroll', handleScroll);
			window.removeEventListener('hashchange', applyHashSection);
			observer.disconnect();
		};
	});
</script>

<svelte:head>
	{#if !isAdmin}
		<title>{siteTitle}</title>
		<meta name="description" content={siteDescription} />
		<meta name="keywords" content="rasas aufar, IT implementation, portfolio, infrastructure, government technology" />
		<link rel="canonical" href="https://rasasaufar.site" />
		<meta property="og:title" content={siteTitle} />
		<meta property="og:description" content={siteDescription} />
		<meta property="og:image" content="https://rasasaufar.site/images/profile.png" />
		<meta property="og:url" content="https://rasasaufar.site" />
		<meta property="og:type" content="website" />
		<meta name="twitter:card" content="summary_large_image" />
		<link rel="icon" href="/favicon.png" />
	{@html `<script type="application/ld+json">${JSON.stringify({
		"@context": "https://schema.org",
		"@type": "Person",
		"name": "Rasas Aufar",
		"url": "https://rasasaufar.site",
		"jobTitle": "IT Implementation Professional",
		"description": "IT Implementation Professional experienced in government technology delivery, infrastructure operations, and applied data analysis.",
		"image": "https://rasasaufar.site/images/profile.png",
		"sameAs": [
			"https://github.com/rasasaufar",
			"https://www.linkedin.com/in/rasasaufar/",
			"https://instagram.com/rasasaufar"			
   			]	
  		})}</script>`}
 	{/if}
</svelte:head>

{#if isAdmin}
	<!-- Admin pages render without public layout -->
	{@render children()}
{:else if isMaintenanceMode}
	<!-- Maintenance Mode -->
	<div class="maintenance-wrap">
		<!-- Animated grid lines -->
		<div class="maintenance-grid-overlay" aria-hidden="true"></div>

		<div class="maintenance-card">
			<div class="maintenance-badge" aria-hidden="true">⚙</div>
			<h1 class="maintenance-title" data-text="System Maintenance">System Maintenance</h1>
			<div class="maintenance-divider"></div>
			<p class="maintenance-subtitle">
				<span class="typewriter">// CURRENTLY OFFLINE FOR UPDATES //</span>
			</p>
			<p class="maintenance-message">
				This portfolio is undergoing scheduled maintenance and improvements. 
				All systems will be back online shortly.
			</p>

			<!-- Progress bar -->
			<div class="maintenance-progress">
				<div class="maintenance-progress-bar"></div>
			</div>
			<p class="maintenance-progress-label">Optimizing systems...</p>

			<div class="maintenance-status">
				<div class="maintenance-status-row">
					<span class="maintenance-dot pulse"></span>
					<span>Maintenance in progress</span>
				</div>
			</div>
			<div class="maintenance-footer">
				<p class="maintenance-terminal">
					<span class="terminal-line">$ sys.status: maintenance</span>
					<span class="terminal-line">$ est. return: soon</span>
					<span class="terminal-line terminal-cursor">$ _</span>
				</p>
			</div>
		</div>

		<!-- Floating decorative elements -->
		<div class="maintenance-decor maintenance-decor-1" aria-hidden="true"></div>
		<div class="maintenance-decor maintenance-decor-2" aria-hidden="true"></div>
		<div class="maintenance-decor maintenance-decor-3" aria-hidden="true"></div>
		<div class="maintenance-decor maintenance-decor-4" aria-hidden="true"></div>
		<div class="maintenance-decor maintenance-decor-5" aria-hidden="true"></div>

		<!-- Scan line effect -->
		<div class="maintenance-scanline" aria-hidden="true"></div>
	</div>
{:else}
	<div class="nav-wrap" class:nav-scrolled={navScrolled}>
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

	<!-- Back to Top Button -->
	<button
		type="button"
		class="back-to-top"
		class:visible={showBackToTop}
		onclick={scrollToTop}
		aria-label="Back to top"
	>
		<svg viewBox="0 0 24 24" aria-hidden="true">
			<polyline points="18 15 12 9 6 15"></polyline>
		</svg>
	</button>
{/if}

<style>
	/* ═══════════════════════════════════════
	   MAINTENANCE MODE
	   ═══════════════════════════════════════ */
	.maintenance-wrap {
		display: flex;
		align-items: center;
		justify-content: center;
		min-height: 100vh;
		padding: 1.5rem;
		position: relative;
		overflow: hidden;
		background:
			radial-gradient(circle at 15% 25%, rgba(255, 102, 145, 0.25) 0 14%, transparent 14.5%),
			radial-gradient(circle at 85% 20%, rgba(117, 255, 191, 0.3) 0 10%, transparent 10.5%),
			radial-gradient(circle at 25% 80%, rgba(110, 174, 255, 0.25) 0 12%, transparent 12.5%),
			radial-gradient(circle at 75% 75%, rgba(248, 228, 92, 0.35) 0 8%, transparent 8.5%),
			#f8e45c;
	}

	/* Animated grid overlay */
	.maintenance-grid-overlay {
		position: absolute;
		inset: 0;
		background:
			repeating-linear-gradient(0deg, rgba(0, 0, 0, 0.03) 0 1px, transparent 1px 35px),
			repeating-linear-gradient(90deg, rgba(0, 0, 0, 0.03) 0 1px, transparent 1px 35px);
		animation: grid-drift 20s linear infinite;
	}

	/* Scan line effect — subtle CRT vibe */
	.maintenance-scanline {
		position: absolute;
		inset: 0;
		background: linear-gradient(
			to bottom,
			transparent 50%,
			rgba(0, 0, 0, 0.015) 50%
		);
		background-size: 100% 4px;
		pointer-events: none;
		z-index: 10;
	}

	/* Card entrance */
	.maintenance-card {
		background: #fffef8;
		border: 3px solid #101010;
		box-shadow: 8px 8px 0 0 #101010;
		padding: 3rem 2.5rem;
		max-width: 520px;
		width: 100%;
		text-align: center;
		position: relative;
		z-index: 2;
		animation: card-entrance 0.8s cubic-bezier(0.16, 1, 0.3, 1) both;
	}

	.maintenance-badge {
		font-size: 3rem;
		margin-bottom: 1rem;
		display: inline-block;
		animation: spin-slow 4s linear infinite;
	}

	/* Glitch title effect */
	.maintenance-title {
		font-family: 'Bebas Neue', sans-serif;
		font-size: clamp(2.5rem, 6vw, 3.5rem);
		letter-spacing: 3px;
		text-transform: uppercase;
		margin: 0;
		line-height: 1;
		color: #101010;
		position: relative;
		animation: title-reveal 1s cubic-bezier(0.16, 1, 0.3, 1) 0.3s both;
	}

	.maintenance-title::before,
	.maintenance-title::after {
		content: attr(data-text);
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		overflow: hidden;
		clip-path: inset(0);
	}

	.maintenance-title::before {
		color: #ff6691;
		animation: glitch-1 4s ease-in-out infinite;
	}

	.maintenance-title::after {
		color: #6eaeff;
		animation: glitch-2 4s ease-in-out infinite;
	}

	.maintenance-divider {
		width: 80px;
		height: 3px;
		background: #101010;
		margin: 1.25rem auto;
		animation: divider-expand 0.6s cubic-bezier(0.16, 1, 0.3, 1) 0.6s both;
		transform-origin: center;
	}

	/* Typewriter effect */
	.maintenance-subtitle {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.72rem;
		text-transform: uppercase;
		letter-spacing: 1.5px;
		color: #505050;
		margin: 0 0 1.5rem;
		animation: fade-up 0.6s ease 0.8s both;
	}

	.typewriter {
		display: inline-block;
		overflow: hidden;
		white-space: nowrap;
		border-right: 2px solid #505050;
		animation: 
			typing 2.5s steps(35) 1.2s both,
			blink-caret 0.75s step-end infinite;
	}

	.maintenance-message {
		font-family: 'Space Grotesk', sans-serif;
		font-size: 1rem;
		line-height: 1.6;
		color: #303030;
		margin: 0 0 1.5rem;
		animation: fade-up 0.6s ease 1s both;
	}

	/* Progress bar */
	.maintenance-progress {
		width: 100%;
		height: 6px;
		background: #e8e8e8;
		border: 2px solid #101010;
		margin-bottom: 0.5rem;
		overflow: hidden;
		animation: fade-up 0.6s ease 1.2s both;
	}

	.maintenance-progress-bar {
		height: 100%;
		background: linear-gradient(90deg, #f8e45c, #d9f7d6, #d4e7ff, #ffd1df, #f8e45c);
		background-size: 200% 100%;
		animation: progress-sweep 2.5s ease-in-out infinite, progress-gradient 3s linear infinite;
	}

	.maintenance-progress-label {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.62rem;
		text-transform: uppercase;
		letter-spacing: 1px;
		color: #505050;
		margin: 0 0 1.5rem;
		animation: fade-up 0.6s ease 1.3s both, text-flicker 3s ease-in-out infinite;
	}

	/* Status badge */
	.maintenance-status {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		background: #d9f7d6;
		border: 2px solid #101010;
		box-shadow: 3px 3px 0 0 #101010;
		padding: 0.6rem 1.2rem;
		margin-bottom: 2rem;
		animation: fade-up 0.6s ease 1.4s both;
		transition: transform 0.2s ease, box-shadow 0.2s ease;
	}

	.maintenance-status:hover {
		transform: translate(-2px, -2px);
		box-shadow: 5px 5px 0 0 #101010;
	}

	.maintenance-status-row {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.72rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
	}

	.maintenance-dot {
		width: 8px;
		height: 8px;
		border-radius: 50%;
		background: #2ecc71;
		border: 1.5px solid #101010;
	}

	.maintenance-dot.pulse {
		animation: pulse-dot 1.5s ease-in-out infinite;
	}

	/* Terminal footer */
	.maintenance-footer {
		border-top: 2px solid #101010;
		padding-top: 1rem;
		animation: fade-up 0.6s ease 1.6s both;
	}

	.maintenance-terminal {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		gap: 0.25rem;
		margin: 0;
	}

	.terminal-line {
		font-family: 'JetBrains Mono', monospace;
		font-size: 0.68rem;
		text-transform: uppercase;
		letter-spacing: 1px;
		color: #505050;
		opacity: 0;
		animation: terminal-type 0.4s ease forwards;
	}

	.terminal-line:nth-child(1) { animation-delay: 2s; }
	.terminal-line:nth-child(2) { animation-delay: 2.5s; }
	.terminal-line:nth-child(3) { animation-delay: 3s; }

	.terminal-cursor {
		animation: terminal-type 0.4s ease 3s forwards, blink-caret 0.75s step-end 3.4s infinite;
		border-right: none;
	}

	/* Decorative elements */
	.maintenance-decor {
		position: absolute;
		border: 3px solid #101010;
		z-index: 1;
	}

	.maintenance-decor-1 {
		width: 50px;
		height: 50px;
		background: #ffd1df;
		box-shadow: 4px 4px 0 0 #101010;
		top: 12%;
		left: 8%;
		animation: float-1 6s ease-in-out infinite, decor-entrance 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.2s both;
	}

	.maintenance-decor-2 {
		width: 35px;
		height: 35px;
		background: #d4e7ff;
		box-shadow: 3px 3px 0 0 #101010;
		bottom: 15%;
		right: 10%;
		animation: float-2 5s ease-in-out infinite, decor-entrance 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.4s both;
	}

	.maintenance-decor-3 {
		width: 28px;
		height: 28px;
		background: #d9f7d6;
		box-shadow: 3px 3px 0 0 #101010;
		top: 18%;
		right: 12%;
		animation: float-3 7s ease-in-out infinite, decor-entrance 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.6s both;
	}

	.maintenance-decor-4 {
		width: 22px;
		height: 22px;
		background: #f8e45c;
		box-shadow: 3px 3px 0 0 #101010;
		bottom: 25%;
		left: 12%;
		border-radius: 50%;
		animation: float-4 8s ease-in-out infinite, decor-entrance 0.8s cubic-bezier(0.16, 1, 0.3, 1) 0.8s both;
	}

	.maintenance-decor-5 {
		width: 18px;
		height: 18px;
		background: #ffd1df;
		box-shadow: 2px 2px 0 0 #101010;
		top: 60%;
		right: 20%;
		animation: float-5 6.5s ease-in-out infinite, decor-entrance 0.8s cubic-bezier(0.16, 1, 0.3, 1) 1s both;
	}

	/* ═══════════════════════════════════════
	   KEYFRAMES
	   ═══════════════════════════════════════ */
	@keyframes card-entrance {
		from {
			opacity: 0;
			transform: translateY(30px) scale(0.96);
		}
		to {
			opacity: 1;
			transform: translateY(0) scale(1);
		}
	}

	@keyframes title-reveal {
		from {
			opacity: 0;
			transform: translateY(15px);
			letter-spacing: 8px;
		}
		to {
			opacity: 1;
			transform: translateY(0);
			letter-spacing: 3px;
		}
	}

	@keyframes divider-expand {
		from { transform: scaleX(0); }
		to { transform: scaleX(1); }
	}

	@keyframes fade-up {
		from {
			opacity: 0;
			transform: translateY(10px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}

	@keyframes typing {
		from { max-width: 0; }
		to { max-width: 100%; }
	}

	@keyframes blink-caret {
		from, to { border-color: transparent; }
		50% { border-color: #505050; }
	}

	@keyframes spin-slow {
		from { transform: rotate(0deg); }
		to { transform: rotate(360deg); }
	}

	@keyframes pulse-dot {
		0%, 100% { opacity: 1; transform: scale(1); }
		50% { opacity: 0.4; transform: scale(1.4); }
	}

	@keyframes progress-sweep {
		0% { width: 15%; }
		50% { width: 85%; }
		100% { width: 15%; }
	}

	@keyframes progress-gradient {
		from { background-position: 0% 0%; }
		to { background-position: 200% 0%; }
	}

	@keyframes text-flicker {
		0%, 100% { opacity: 1; }
		92% { opacity: 1; }
		93% { opacity: 0.4; }
		94% { opacity: 1; }
		96% { opacity: 0.4; }
		97% { opacity: 1; }
	}

	@keyframes terminal-type {
		from { opacity: 0; transform: translateX(-5px); }
		to { opacity: 1; transform: translateX(0); }
	}

	@keyframes glitch-1 {
		0%, 100% { clip-path: inset(0 0 100% 0); }
		4% { clip-path: inset(20% 0 60% 0); transform: translateX(-2px); }
		5% { clip-path: inset(0 0 100% 0); transform: translateX(0); }
		45% { clip-path: inset(0 0 100% 0); }
		46% { clip-path: inset(60% 0 10% 0); transform: translateX(2px); }
		47% { clip-path: inset(0 0 100% 0); transform: translateX(0); }
	}

	@keyframes glitch-2 {
		0%, 100% { clip-path: inset(0 0 100% 0); }
		20% { clip-path: inset(0 0 100% 0); }
		21% { clip-path: inset(40% 0 30% 0); transform: translateX(2px); }
		22% { clip-path: inset(0 0 100% 0); transform: translateX(0); }
		70% { clip-path: inset(0 0 100% 0); }
		71% { clip-path: inset(70% 0 5% 0); transform: translateX(-2px); }
		72% { clip-path: inset(0 0 100% 0); transform: translateX(0); }
	}

	@keyframes grid-drift {
		from { transform: translate(0, 0); }
		to { transform: translate(35px, 35px); }
	}

	@keyframes float-1 {
		0%, 100% { transform: rotate(12deg) translate(0, 0); }
		50% { transform: rotate(18deg) translate(8px, -12px); }
	}

	@keyframes float-2 {
		0%, 100% { transform: rotate(-8deg) translate(0, 0); }
		50% { transform: rotate(-14deg) translate(-6px, 10px); }
	}

	@keyframes float-3 {
		0%, 100% { transform: rotate(25deg) translate(0, 0); }
		50% { transform: rotate(30deg) translate(5px, -8px); }
	}

	@keyframes float-4 {
		0%, 100% { transform: translate(0, 0) scale(1); }
		50% { transform: translate(-5px, 8px) scale(1.1); }
	}

	@keyframes float-5 {
		0%, 100% { transform: rotate(0deg) translate(0, 0); }
		50% { transform: rotate(10deg) translate(6px, -6px); }
	}

	@keyframes decor-entrance {
		from {
			opacity: 0;
			transform: scale(0) rotate(0deg);
		}
		to {
			opacity: 1;
			transform: scale(1) rotate(var(--rotate, 0deg));
		}
	}

	/* ═══════════════════════════════════════
	   RESPONSIVE
	   ═══════════════════════════════════════ */
	@media (max-width: 600px) {
		.maintenance-card {
			padding: 2rem 1.5rem;
		}

		.maintenance-decor-1 {
			width: 35px;
			height: 35px;
		}

		.maintenance-decor-2 {
			width: 25px;
			height: 25px;
		}

		.maintenance-decor-3,
		.maintenance-decor-4,
		.maintenance-decor-5 {
			display: none;
		}

		.typewriter {
			white-space: normal;
			border-right: none;
			animation: fade-up 0.6s ease 1.2s both;
		}
	}

	/* Respect reduced motion */
	@media (prefers-reduced-motion: reduce) {
		.maintenance-badge,
		.maintenance-decor,
		.maintenance-progress-bar,
		.maintenance-dot,
		.maintenance-grid-overlay {
			animation: none;
		}

		.maintenance-title::before,
		.maintenance-title::after {
			display: none;
		}

		.typewriter {
			animation: fade-up 0.6s ease both;
			border-right: none;
		}

		.maintenance-card {
			animation: fade-up 0.5s ease both;
		}

		.terminal-line {
			opacity: 1;
			animation: none;
		}
	}
</style>
