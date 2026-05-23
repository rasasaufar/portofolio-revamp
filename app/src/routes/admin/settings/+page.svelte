<script lang="ts">
	import { onMount } from 'svelte';
	import { listResource, updateResource } from '$lib/api/***REMOVED***';

	let formData = $state<Record<string, unknown>>({});
	let loading = $state(true);
	let saving = $state(false);
	let toast = $state('');

	onMount(async () => {
		try {
			const items = await listResource('settings');
			if (items.length > 0) formData = items[0];
		} catch {}
		loading = false;
	});

	async function handleSave(e: SubmitEvent) {
		e.preventDefault();
		saving = true;
		try {
			await updateResource('settings', String(formData.id), {
				site_title: formData.site_title,
				meta_description: formData.meta_description,
				favicon_url: formData.favicon_url,
				logo_url: formData.logo_url,
				footer_text: formData.footer_text,
				theme_mode: formData.theme_mode,
				maintenance_mode: formData.maintenance_mode
			});
			toast = 'Settings saved!';
			setTimeout(() => { toast = ''; }, 3000);
		} catch (e) {
			toast = 'Error: ' + (e instanceof Error ? e.message : 'Save failed');
		}
		saving = false;
	}
</script>

<div class="***REMOVED***-header"><h1>Site Settings</h1></div>

{#if loading}
	<p style="color: var(--***REMOVED***-text-muted);">Loading...</p>
{:else}
	<div class="***REMOVED***-card">
		<form class="***REMOVED***-form" onsubmit={handleSave}>
			<div class="***REMOVED***-field"><label for="site_title">Site Title</label><input id="site_title" class="***REMOVED***-input" bind:value={formData.site_title as string} /></div>
			<div class="***REMOVED***-field"><label for="meta_description">Meta Description</label><textarea id="meta_description" class="***REMOVED***-input ***REMOVED***-textarea" bind:value={formData.meta_description as string}></textarea></div>
			<div class="***REMOVED***-field"><label for="favicon_url">Favicon URL</label><input id="favicon_url" class="***REMOVED***-input" bind:value={formData.favicon_url as string} /></div>
			<div class="***REMOVED***-field"><label for="logo_url">Logo URL</label><input id="logo_url" class="***REMOVED***-input" bind:value={formData.logo_url as string} /></div>
			<div class="***REMOVED***-field"><label for="footer_text">Footer Text</label><input id="footer_text" class="***REMOVED***-input" bind:value={formData.footer_text as string} /></div>
			<div class="***REMOVED***-field"><label for="theme_mode">Theme Mode</label><select id="theme_mode" class="***REMOVED***-input" bind:value={formData.theme_mode as string}><option value="dark">Dark</option><option value="light">Light</option></select></div>
			<div class="***REMOVED***-field"><label style="display:flex;align-items:center;gap:0.5rem;cursor:pointer;"><input type="checkbox" bind:checked={formData.maintenance_mode as boolean} /> Maintenance Mode</label></div>
			<button class="***REMOVED***-btn ***REMOVED***-btn-primary" type="submit" disabled={saving}>{saving ? 'Saving...' : 'Save Settings'}</button>
		</form>
	</div>
{/if}

{#if toast}<div class="***REMOVED***-toast ***REMOVED***-toast-success">{toast}</div>{/if}
