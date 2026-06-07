package seeds

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/portfolio/api/internal/service"
)

// Run executes the database seed, populating all tables with initial data.
// It is idempotent — skips insertion if data already exists.
func Run(pool *pgxpool.Pool, adminPassword string) error {
	ctx := context.Background()

	log.Println("🌱 Starting database seed...")

	if err := seedAdminUser(ctx, pool, adminPassword); err != nil {
		return fmt.Errorf("seed admin user: %w", err)
	}
	if err := seedIdentityConsole(ctx, pool); err != nil {
		return fmt.Errorf("seed identity console: %w", err)
	}
	if err := seedCapabilitySnapshots(ctx, pool); err != nil {
		return fmt.Errorf("seed capability snapshots: %w", err)
	}
	if err := seedImplementationStrengths(ctx, pool); err != nil {
		return fmt.Errorf("seed implementation strengths: %w", err)
	}
	if err := seedProfessionalDossier(ctx, pool); err != nil {
		return fmt.Errorf("seed professional dossier: %w", err)
	}
	if err := seedEducation(ctx, pool); err != nil {
		return fmt.Errorf("seed education: %w", err)
	}
	if err := seedWorkExperiences(ctx, pool); err != nil {
		return fmt.Errorf("seed work experiences: %w", err)
	}
	if err := seedProjects(ctx, pool); err != nil {
		return fmt.Errorf("seed projects: %w", err)
	}
	if err := seedCertifications(ctx, pool); err != nil {
		return fmt.Errorf("seed certifications: %w", err)
	}
	if err := seedPublications(ctx, pool); err != nil {
		return fmt.Errorf("seed publications: %w", err)
	}
	if err := seedContactInfo(ctx, pool); err != nil {
		return fmt.Errorf("seed contact info: %w", err)
	}
	if err := seedSiteSettings(ctx, pool); err != nil {
		return fmt.Errorf("seed site settings: %w", err)
	}

	log.Println("✅ Database seed completed successfully")
	return nil
}

// Helper to check if a table already has data
func tableHasData(ctx context.Context, pool *pgxpool.Pool, table string) bool {
	var count int
	err := pool.QueryRow(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// Helper to marshal JSON
func mustJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func seedAdminUser(ctx context.Context, pool *pgxpool.Pool, password string) error {
	if tableHasData(ctx, pool, "admin_users") {
		log.Println("  ⏭️  admin_users already seeded, skipping")
		return nil
	}

	authSvc := service.NewAuthService("", 0)
	hash, err := authSvc.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx,
		`INSERT INTO admin_users (email, password_hash, name) VALUES ($1, $2, $3)`,
		"***REMOVED***", hash, "Admin",
	)
	if err != nil {
		return err
	}
	log.Println("  ✅ admin_users seeded")
	return nil
}

func seedIdentityConsole(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "identity_console") {
		log.Println("  ⏭️  identity_console already seeded, skipping")
		return nil
	}

	_, err := pool.Exec(ctx,
		`INSERT INTO identity_console (name, role, headline, description, avatar_url, current_focus, availability_text, cta_primary_label, cta_primary_link, cta_secondary_label, cta_secondary_link, order_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		"Rasas Aufar",
		"IT Implementation Professional Staff",
		"IT Implementation Professional experienced in managing and delivering technology solutions for government projects, ensuring reliable system implementation, process optimization, and alignment with organizational and public sector needs.",
		"IT Implementation Professional experienced in managing and delivering technology solutions for government projects, ensuring reliable system implementation, process optimization, and alignment with organizational and public sector needs.",
		"",
		mustJSON([]string{"Linux", "UI / UX Design", "Technical Documentation", "Database", "Prototyping", "Figma", "Deployment", "CI / CD", "Github Actions", "Data Analysis"}),
		"Available for Opportunities",
		"View Projects",
		"#laboratory",
		"Contact Me",
		"#contact",
		0,
	)
	if err != nil {
		return err
	}
	log.Println("  ✅ identity_console seeded")
	return nil
}

func seedCapabilitySnapshots(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "capability_snapshots") {
		log.Println("  ⏭️  capability_snapshots already seeded, skipping")
		return nil
	}

	stats := []struct {
		Label string
		Value string
		Order int
	}{
		{"Years Exp.", "2+", 0},
		{"Projects", "5+", 1},
		{"Publications", "1", 2},
	}

	for _, s := range stats {
		_, err := pool.Exec(ctx,
			`INSERT INTO capability_snapshots (label, value, order_number) VALUES ($1, $2, $3)`,
			s.Label, s.Value, s.Order,
		)
		if err != nil {
			return err
		}
	}
	log.Println("  ✅ capability_snapshots seeded")
	return nil
}

func seedImplementationStrengths(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "implementation_strengths") {
		log.Println("  ⏭️  implementation_strengths already seeded, skipping")
		return nil
	}

	strengths := []struct {
		Title  string
		Points []string
		Order  int
	}{
		{
			"Infrastructure Reliability",
			[]string{"Ubuntu server administration", "Production database backup and recovery", "Deployment and post-deployment validation"},
			0,
		},
		{
			"Implementation Delivery",
			[]string{"Onsite government system implementation", "Cross-unit stakeholder coordination", "Monthly technical reporting and change requests"},
			1,
		},
		{
			"Training and Enablement",
			[]string{"Technical guidance sessions", "Presenter for defense-sector application training", "Operational handover and user adoption support"},
			2,
		},
	}

	for _, s := range strengths {
		_, err := pool.Exec(ctx,
			`INSERT INTO implementation_strengths (title, bullet_points, order_number) VALUES ($1, $2, $3)`,
			s.Title, mustJSON(s.Points), s.Order,
		)
		if err != nil {
			return err
		}
	}
	log.Println("  ✅ implementation_strengths seeded")
	return nil
}

func seedProfessionalDossier(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "professional_dossier") {
		log.Println("  ⏭️  professional_dossier already seeded, skipping")
		return nil
	}

	_, err := pool.Exec(ctx,
		`INSERT INTO professional_dossier (title, paragraph_1, paragraph_2, paragraph_3) VALUES ($1, $2, $3, $4)`,
		"Professional Dossier",
		"A Computer Science graduate currently pursuing a career as an IT Implementor with specialized capabilities in IT Infrastructure and Data Analysis. Possesses a robust technical foundation in integrating technology solutions, ensuring systems align with operational requirements, and leveraging data-driven insights to support strategic decision-making.",
		"Demonstrates a solid track record in managing large-scale projects within the public sector and government institutions. Experience spans across Linux-based server infrastructure management (Ubuntu), complex database operations, and comprehensive application deployment and functional testing. Proven ability to oversee the end-to-end implementation of information systems, including the management of digital attendance systems and institutional learning platforms.",
		"Beyond technical proficiency, exhibits strong managerial skills through the systematic and accurate preparation of monthly technical documentation. Experienced in leading technical guidance sessions and user training programs, ensuring seamless technology transfer and ensuring that applications are utilized optimally by all stakeholders.",
	)
	if err != nil {
		return err
	}
	log.Println("  ✅ professional_dossier seeded")
	return nil
}

func seedEducation(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "education") {
		log.Println("  ⏭️  education already seeded, skipping")
		return nil
	}

	educations := []struct {
		Institution string
		Degree      string
		Major       string
		StartYear   string
		EndYear     string
		GPA         string
		Description string
		ImageURL    string
		Tags        []string
		Order       int
	}{
		{
			"Telkom University Purwokerto",
			"Bachelor of Informatics Engineering",
			"Informatics Engineering",
			"2018", "2022",
			"3.23",
			"An ordinary college student who was quite active in university sports, particularly Badminton and Futsal, and managed to successfully graduate on time.",
			"/images/education/telkom-university-logo.webp",
			[]string{"Graduated on Time"},
			0,
		},
		{
			"SMAN 4 Kota Pekalongan",
			"Senior High School - Science",
			"Science",
			"2015", "2018",
			"92.5",
			"Excelled in mathematics and science subjects. Participated in national olympiads and coding competitions.",
			"/images/education/sman4-pekalongan-logo.jpg",
			[]string{"National Science Olympiad", "Student Council", "Top 10 Graduate"},
			1,
		},
	}

	for _, e := range educations {
		_, err := pool.Exec(ctx,
			`INSERT INTO education (institution_name, degree, major, start_year, end_year, gpa, description, image_url, tags, order_number)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			e.Institution, e.Degree, e.Major, e.StartYear, e.EndYear, e.GPA, e.Description, e.ImageURL, mustJSON(e.Tags), e.Order,
		)
		if err != nil {
			return err
		}
	}
	log.Println("  ✅ education seeded")
	return nil
}

func seedWorkExperiences(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "work_experiences") {
		log.Println("  ⏭️  work_experiences already seeded, skipping")
		return nil
	}

	type GalleryItem struct {
		Image       string `json:"image"`
		Caption     string `json:"caption"`
		Description string `json:"description"`
	}

	experiences := []struct {
		Company      string
		Position     string
		StartDate    string
		EndDate      string
		IsCurrent    bool
		Description  string
		BulletPoints []string
		TechTags     []string
		LogoURL      string
		Gallery      []GalleryItem
		Order        int
	}{
		{
			"PT. Traspac Makmur Sejahtera",
			"IT Implementation",
			"Aug 2024", "Present",
			true,
			"Serve as an IT Implementor, IT Infrastructure specialist, and Data Analyst, leading multiple projects across government ministries and military institutions.",
			[]string{
				"Kemenlu RI - Worked onsite with the SISTIK Sub-Directorate of the PWNI Directorate to provide technical support for developing the \"Safe Travel\" website and managing the \"Portal Peduli WNI\" server infrastructure using Ubuntu",
				"Kemenlu RI - Prepared monthly technical documentation, including maintenance reports, bug analyses, troubleshooting guides, and change request management",
				"Mabes TNI - Managed critical database operations, including routine backup and recovery in production environments to ensure data integrity",
				"Mabes TNI - Led Sisfopajak application Deployment on physical server infrastructure and executed comprehensive post-deployment functional testing and verification",
				"Mabes TNI - Served as presenter for training on the Alutsista Data System and supporting applications at Mabes TNI, Mabes AL, and the Ministry of Defense Data and Information Center",
				"Ministry of Downstreaming and Investment - Managed the attendance system for BKPM human resources operations",
			},
			[]string{"Ubuntu", "Server Infrastructure", "Database Management", "Application Deployment"},
			"https://karir.traspac.id/assets/img/logotraspac.png",
			[]GalleryItem{
				{"/images/gallery/traspac-kemenlu.png", "Ministry of Foreign Affairs of the Republic of Indonesia", "Delivered technical support and technical guidance activities with the SISTIK Sub-Directorate of the PWNI Directorate at the Ministry of Foreign Affairs of the Republic of Indonesia."},
				{"/images/gallery/traspac-kemhan.png", "Ministry of Defense Data and Information Center", "Served as the training presenter for the Alutsista Data System and supporting applications at the Ministry of Defense Data and Information Center of the Republic of Indonesia."},
				{"/images/gallery/traspac-mabesal.png", "Mabes TNI AL - Gedung Denma", "Presented the Alutsista Data System application for the Indonesian Navy (TNI AL) and served as a trainer for the application simulation session at Gedung Denma Mabes AL, with an audience of approximately 200 personnel."},
			},
			0,
		},
		{
			"PT. Metro Network Solutions",
			"Operator Scanner",
			"Sep 2023", "Dec 2023",
			false,
			"Served on the Archive Digitization project team at the National Land Agency (BPN) of Pekalongan City, supporting the transformation of physical records into digital archives.",
			[]string{
				"Contributed to the digitization project at the National Land Agency (BPN) Office of Pekalongan City",
				"Sorted and organized physical albums and documents with high accuracy prior to digitization",
				"Performed high-volume daily document scanning to support archival preparation",
				"Edited and optimized scanned outputs using NAPS2 to ensure clarity, alignment, and readability",
				"Maintained high discipline and performance under pressure to consistently meet scanning targets and deadlines",
				"Safeguarded the completeness, confidentiality, and integrity of high-value land documents throughout the project",
			},
			[]string{"NAPS2", "Document Editing", "Data Management", "Scanner Handling"},
			"https://media.licdn.com/dms/image/v2/C560BAQHq1FwKfCDJZA/company-logo_200_200/company-logo_200_200/0/1667985560103/pt_metronet_logo?e=2147483647&v=beta&t=gsRuAkezzvj65j2AGShob4KmQDCnAjpuGbO4YH5xGZE",
			[]GalleryItem{
				{"/images/gallery/bpn-closing.jpeg", "Project Closing", "Team closing moment with the full implementation team at the end of the BPN Pekalongan City Archive Digitization project after successfully achieving all scanning targets."},
			},
			1,
		},
		{
			"Niagahoster - Web Hosting Unlimited Indonesia",
			"Project-Based Virtual Intern : UI / UX Designer",
			"Jul 2023", "Aug 2023",
			false,
			"Completed a project-based virtual internship with Rakamin Academy (Niagahoster x Rakamin Academy).",
			[]string{
				"Completed project assignments aligned with the UI/UX Designer role at Niagahoster",
				"Designed user interfaces and user experiences to improve usability and customer flow",
				"Built wireframes and interactive prototypes to visualize end-to-end user journeys, including Checkout Flow",
			},
			[]string{"User Experience (UX)", "Prototyping", "User Interface", "Wireframing"},
			"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRVYFavWOvYAbq_z018Y5-cZNuutUC_256vVw&s",
			[]GalleryItem{
				{"/images/gallery/Niagahoster.png", "Checkout Flow", "Redesigned Checkout Flow for the Niagahoster website to deliver a secure, intuitive, and efficient shopping and payment experience."},
				{"/images/gallery/5.png", "Certificate of Completion", "Official completion certificate validating achievements, successful program delivery, and essential UI/UX skill development during the virtual internship."},
			},
			2,
		},
		{
			"Nuri",
			"Project-Based Virtual Intern : UI / UX Designer",
			"Jul 2023", "Aug 2023",
			false,
			"Completed a project-based virtual internship with Rakamin Academy (Nuri x Rakamin Academy).",
			[]string{
				"Completed project assignments aligned with the UI/UX Designer role at Nuri",
				"Developed design solutions from research through User Interface and User Experience execution",
				"Created design deliverables from low-fidelity wireframes to high-fidelity prototypes",
			},
			[]string{"User Experience (UX)", "Prototyping", "User Interface", "Wireframing"},
			"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRxW4zjU4gLFxMQ7K2mBzwpxOHJId1nbb28LA&s",
			[]GalleryItem{
				{"/images/gallery/Nuri.png", "Nuri Final Task", "Final project focused on designing the user interface for the COD feature on the Nuri e-commerce platform."},
				{"/images/gallery/6.png", "Certificate of Completion", "Official completion certificate from the Nuri x Rakamin Academy virtual internship program."},
			},
			3,
		},
		{
			"Dinas Komunikasi dan Informatika Kab. Pekalongan",
			"Student Internship",
			"Sep 2021", "Oct 2021",
			false,
			"Completed an internship focused on interface design and supporting system development for local government initiatives.",
			[]string{
				"Designed the UI/UX for an e-commerce application aimed at supporting UMKM growth in Pekalongan Regency",
				"Rapidly acquired new skills and applied them to daily tasks to improve operational efficiency and productivity",
				"Collaborated actively in the agency work environment, supported team members, and aligned deliverables with user needs",
			},
			[]string{"UI/UX Design", "Figma", "User Research", "Prototyping"},
			"https://upload.wikimedia.org/wikipedia/commons/7/74/Lambang_Kabupaten_Pekalongan.JPG",
			[]GalleryItem{
				{"/images/gallery/dinkominfo.jpg", "Dinkominfo Office", "Office of the Pekalongan Regency Department of Communication and Informatics, where the internship program was carried out."},
				{"/images/gallery/Magang.png", "E-Commerce UMKM", "UI/UX design for an e-commerce platform developed to be accessible and easy to use for UMKM business owners."},
			},
			4,
		},
	}

	for _, e := range experiences {
		_, err := pool.Exec(ctx,
			`INSERT INTO work_experiences (company_name, position, start_date, end_date, is_current, description, bullet_points, tech_tags, logo_url, gallery_images, order_number)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			e.Company, e.Position, e.StartDate, e.EndDate, e.IsCurrent, e.Description,
			mustJSON(e.BulletPoints), mustJSON(e.TechTags), e.LogoURL, mustJSON(e.Gallery), e.Order,
		)
		if err != nil {
			return err
		}
	}
	log.Println("  ✅ work_experiences seeded")
	return nil
}

func seedProjects(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "projects") {
		log.Println("  ⏭️  projects already seeded, skipping")
		return nil
	}

	projects := []struct {
		Title      string
		Category   string
		Desc       string
		TechTags   []string
		ImageURL   string
		DemoURL    string
		RepoURL    string
		IsFeatured bool
		Order      int
	}{
		{"Automated Database Backup & Rolling Retention System", "Infrastructure", "Engineered a robust Bash scripting solution for automated daily backups via Cron. Implemented a 10-day rolling retention policy across dual servers to optimize storage and ensure disaster recovery readiness. Managed application deployments utilizing SVN.", []string{"Linux", "Bash", "Cron", "SVN", "Disaster Recovery"}, "/images/portfolio/automated-db-backup-rolling-retention.png", "", "", true, 0},
		{"Server Deployment & Process Management", "Infrastructure", "Executed the end-to-end deployment of the SISFOPAJAK application on bare-metal Ubuntu servers. Configured process management using PM2 to ensure zero-downtime restarts and high availability in a high-security production environment.", []string{"Ubuntu", "PM2", "Deployment", "Server Admin"}, "/images/portfolio/server-deployment-pm2.png", "", "", true, 1},
		{"Multi-App Containerization & Docker Orchestration", "Infrastructure", "Migrated web application stack from PM2 process management to Docker containerization. Orchestrated multiple SvelteKit applications with docker-compose, implementing isolated environments, automated restarts, and consistent deployment pipelines across projects.", []string{"Docker", "docker-compose", "Containerization", "SvelteKit", "Migration"}, "/images/portfolio/docker-orchestration.png", "", "", true, 2},
		{"Multi-Domain Reverse Proxy & SSL Infrastructure", "Infrastructure", "Designed and configured Nginx reverse proxy infrastructure to serve multiple web applications across subdomains from a single OCI instance. Automated SSL certificate provisioning and renewal using Certbot/Let's Encrypt, enforcing HTTPS-only access with proper header forwarding for WebSocket support.", []string{"Nginx", "Reverse Proxy", "Let's Encrypt", "SSL", "Multi-tenant"}, "/images/portfolio/nginx-ssl-proxy.png", "", "", true, 3},
		{"Server Security Hardening & Intrusion Prevention", "Infrastructure", "Implemented multi-layered server security posture on production OCI instance. Configured Fail2Ban with custom email notification pipeline (msmtp/Gmail), enforced SSH key-only authentication, applied OCI Security List ingress rules, and established IP whitelisting policy.", []string{"Fail2Ban", "SSH Hardening", "Firewall", "OCI Security", "msmtp"}, "/images/portfolio/server-security-hardening.png", "", "", true, 4},
		{"Production Server Monitoring & Dashboard System", "Infrastructure", "Built automated server health monitoring system with terminal dashboard (Figlet/ASCII art) for real-time system metrics visibility. Implemented scheduled reporting pipeline delivering daily infrastructure status summaries to team communication channels via cron automation.", []string{"Monitoring", "Cron", "Automation", "Discord Integration", "System Admin"}, "/images/portfolio/server-monitoring-dashboard.png", "", "", false, 5},
		{"Full-Stack Deployment Pipeline (Laravel API + Nuxt Frontend)", "Infrastructure", "Deployed full-stack application architecture on single VPS instance — Laravel REST API backend with Nuxt.js SSR frontend. Configured Nginx location-based routing to proxy API, static assets, and frontend through unified domain entry point with proper header forwarding.", []string{"Laravel", "Nuxt.js", "Nginx", "VPS", "Full-Stack"}, "/images/portfolio/fullstack-laravel-nuxt.png", "", "", false, 6},
		{"CI/CD Pipeline (GitHub Actions)", "Infrastructure", "Designed and implemented a GitOps-style CI/CD pipeline using GitHub Actions for automated production deployment. On every push to the main branch, the pipeline establishes a secure SSH connection to the OCI VPS via encrypted repository secrets, pulls the latest source code, and triggers a zero-downtime Docker Compose rebuild and rollout.", []string{"GitHub Actions", "CI/CD", "Docker", "SSH", "GitOps", "Automation"}, "/images/portfolio/cicd-github-actions.png", "", "", false, 7},
		{"Final Project Redesign Maxim", "UI/UX", "Redesigned the Maxim application and added Top Up and payment features.", []string{"UI/UX Design", "Figma", "Prototyping"}, "/images/portfolio/Maxim.png", "https://drive.google.com/file/d/14rentR_FyblFyTi5F63m0UeVRDSomjyJ/view?usp=sharing", "", false, 8},
		{"Final Task Niagahoster x Rakamin Academy", "UI/UX", "Designed a Checkout Flow for the Niagahoster website.", []string{"UI/UX Design", "Figma"}, "/images/portfolio/Niagahoster.png", "https://drive.google.com/file/d/1GdUdGYV2Tg1zPnZXwADqJSzOUguJFTPj/view?usp=sharing", "", false, 9},
		{"Final Task Nuri x Rakamin Academy", "UI/UX", "Designed and integrated a Cash On Delivery (COD) feature for an e-commerce application.", []string{"UI/UX Design", "Figma"}, "/images/portfolio/Nuri.png", "https://drive.google.com/file/d/1GoF12FssWWEjVzu2kExP58ZMcDjMV4zG/view?usp=sharing", "", false, 10},
		{"UMKM Shopping Application in Kajen", "UI/UX", "Designed the UI for a shopping application for UMKM in Kajen, Pekalongan Regency.", []string{"UI/UX Design", "Figma"}, "/images/portfolio/Magang.png", "https://bit.ly/ProtoypeMagang", "", false, 11},
		{"Fake Project", "UI/UX", "Designed and developed a restaurant website landing page.", []string{"UI/UX Design", "Figma"}, "/images/portfolio/Aarss.png", "https://drive.google.com/file/d/1hqhdVyanQeePJewHR8PJ9kK07BdrkfOJ/view?usp=sharing", "", false, 12},
		{"Mini Task in Short Class @myskill.id", "UI/UX", "Completed the design tasks assigned during the mini task program.", []string{"UI/UX Design", "Figma"}, "/images/portfolio/Minitask.png", "https://www.figma.com/proto/LsBpusfnm1lkL8gUjjGSLR/MiniTask---MySkill?page-id=0%3A1&type=design&node-id=2-541&viewport=609%2C593%2C0.36&t=Fq9mTWEDSw0SxcEg-1&scaling=scale-down&starting-point-node-id=2%3A541&mode=design", "", false, 13},
	}

	for _, p := range projects {
		_, err := pool.Exec(ctx,
			`INSERT INTO projects (title, category, description, tech_tags, image_url, demo_url, repo_url, is_featured, order_number)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			p.Title, p.Category, p.Desc, mustJSON(p.TechTags), p.ImageURL, p.DemoURL, p.RepoURL, p.IsFeatured, p.Order,
		)
		if err != nil {
			return err
		}
	}
	log.Println("  ✅ projects seeded")
	return nil
}

func seedCertifications(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "certifications") {
		log.Println("  ⏭️  certifications already seeded, skipping")
		return nil
	}

	certs := []struct {
		Title        string
		Issuer       string
		IssuedDate   string
		CredentialID string
		Description  string
		Skills       []string
		ImageURL     string
		Category     string
		Order        int
	}{
		{"Wawasan Karir dalam Bidang Data Analytics", "Digital Talent Scholarship", "Jul 2025", "2299746850-4500", "Career insights training in Data Analytics delivered through the Digital Talent Scholarship program.", []string{"Data Analytics", "Data Visualization"}, "/images/certificates/dts-logo.png", "additional", 0},
		{"Pengenalan Data Science dan Pemanfaatannya dalam Berbagai Sektor", "Digital Talent Scholarship", "Jul 2025", "2299818850-4751", "Foundational Data Science training and cross-sector application through the Digital Talent Scholarship program.", []string{"Data Science", "Data Analysis"}, "/images/certificates/dts-logo.png", "additional", 1},
		{"Keamanan IT: Pertahanan terhadap Kejahatan Digital", "Coursera", "Jul 2024", "TDN7P3MAAEG7", "IT security course under the Google IT Support Specialization.", []string{"Cybersecurity", "Encryption", "Network Security"}, "/images/certificates/coursera-logo.png", "core", 2},
		{"Administrasi Sistem dan Layanan Infrastruktur TI", "Coursera", "Jul 2024", "RC8NK9PBR5XG", "System administration and IT infrastructure services course under the Google IT Support Specialization.", []string{"Server Management", "Directory Services", "IT Infrastructure"}, "/images/certificates/coursera-logo.png", "core", 3},
		{"Sistem Operasi dan Anda: Menjadi Pengguna yang Berdaya", "Coursera", "Jul 2024", "7KU8JRMPQVHJ", "Operating systems course under the Google IT Support Specialization.", []string{"Windows", "Linux", "File Management", "Command Line"}, "/images/certificates/coursera-logo.png", "core", 4},
		{"Spesialisasi IT Support Google", "Google", "Jul 2024", "YL552W7AGU2G", "Professional certification from Google covering core IT support competencies.", []string{"Cloud Computing", "Computer Networking", "IT Security", "System Administration", "Troubleshooting", "Operating Systems"}, "/images/certificates/google-logo.png", "core", 5},
		{"Seluk Beluk Jaringan Komputer", "Coursera", "Jun 2024", "2FLPLT9K6MXS", "Computer networking course under the Google IT Support Specialization.", []string{"TCP/IP", "DNS", "DHCP", "Network Troubleshooting"}, "/images/certificates/coursera-logo.png", "core", 6},
		{"Dasar-Dasar Dukungan Teknis", "Coursera", "May 2024", "7C2ZULPDA5P6", "Foundational technical support course under the Google IT Support Specialization.", []string{"Troubleshooting", "Customer Service", "Operating Systems"}, "/images/certificates/coursera-logo.png", "core", 7},
		{"TOEFL", "Englishvit", "Oct 2023", "EV/TO4/10/2023/0159", "English proficiency certificate based on an online TOEFL test, valid through October 2025.", []string{"English Proficiency", "TOEFL"}, "/images/certificates/englishvit-logo.png", "additional", 8},
		{"Certificate of Mastery UI-UX Research and Design: Fullstack Intensive Bootcamp", "MySkill", "Sep 2023", "MS-1/9/2023-5qlzkE6XeBcbQREEVF0I", "Completion certificate for an intensive full-stack UI/UX Research and Design bootcamp.", []string{"User Interface", "User Experience (UX)", "Design Research"}, "/images/certificates/myskill-logo.png", "additional", 9},
		{"Certificate of Competencies - Nuri UI/UX Designer Virtual Internship", "Nuri", "Jul 2023", "212536IAPMGIN2672023", "Competency certificate for the UI/UX Designer role through the Nuri x Rakamin Academy virtual internship program.", []string{"Sketching", "Design Thinking", "User Interface", "User Experience", "Prototyping"}, "/images/certificates/nuri-logo.png", "additional", 10},
		{"Figma For UI/UX Design", "MySkill", "Jul 2023", "48508/UIX/LM/07/2023", "UI/UX design training using Figma.", []string{"Figma", "User Interface Design"}, "/images/certificates/myskill-logo.png", "additional", 11},
		{"Fundamental UX Design", "Coding Studio", "Jul 2023", "77DF6D121A-77F17BB889-75BB1E65BD", "Fundamental User Experience Design training.", []string{"User Experience (UX)", "UX Research", "UX Writing"}, "/images/certificates/coding-studio-logo.png", "additional", 12},
		{"Fundamental UI Design", "Coding Studio", "Jul 2023", "77DF6D121A-75C17F8C7B-75BB1E65BD", "Fundamental User Interface Design training.", []string{"User Interface Design", "Visual Design", "Layout"}, "/images/certificates/coding-studio-logo.png", "additional", 13},
		{"Certificate of Competencies - Niagahoster UI/UX Designer Virtual Internship", "Niagahoster - Web Hosting Unlimited Indonesia", "Jul 2023", "212536IAPMGIN3072023", "Competency certificate for the UI/UX Designer role through the Niagahoster x Rakamin Academy virtual internship program.", []string{"User Journeys", "Process Design", "UI Design", "UX Research", "Wireframing", "Prototyping"}, "/images/certificates/niagahoster-logo.png", "additional", 14},
		{"AWS Certified Cloud Practitioner", "Amazon Web Services (AWS)", "May 2023", "", "Official AWS certification validating foundational knowledge of AWS Cloud services, architecture, and best practices, valid through May 2026.", []string{"Cloud Computing", "AWS Services", "Cloud Architecture"}, "/images/certificates/aws-cloud-practitioner.png", "core", 15},
		{"AWS Cloud Foundation", "Digital Talent Scholarship", "May 2023", "1955634840-1099/FGA/BLSDM.Kominfo/2023", "Training certificate awarded through the Digital Talent Scholarship program by Kominfo.", []string{"Cloud Computing"}, "/images/certificates/dts-logo.png", "core", 16},
		{"AWS Academy Graduate - AWS Academy Cloud Foundations", "Amazon Web Services (AWS)", "Mar 2023", "", "Completion of the AWS Academy Cloud Foundations program covering cloud fundamentals, core AWS services, security, and architecture.", []string{"Cloud Foundations", "AWS Core Services", "Cloud Security"}, "/images/certificates/aws-cloud-security.png", "core", 17},
		{"Certificate Of Competence Junior Web Developer", "Badan Nasional Sertifikasi Profesi (BNSP)", "Aug 2022", "62019 2513 18247 2022", "National competency certification as a Junior Web Developer issued by BNSP, valid through August 2025.", []string{"CSS", "HTML", "Web Development"}, "/images/certificates/bnsp-logo.png", "core", 18},
	}

	for _, c := range certs {
		_, err := pool.Exec(ctx,
			`INSERT INTO certifications (title, issuer, issued_date, credential_id, description, skills, image_url, category, order_number)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
			c.Title, c.Issuer, c.IssuedDate, c.CredentialID, c.Description, mustJSON(c.Skills), c.ImageURL, c.Category, c.Order,
		)
		if err != nil {
			return err
		}
	}
	log.Println("  ✅ certifications seeded")
	return nil
}

func seedPublications(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "publications") {
		log.Println("  ⏭️  publications already seeded, skipping")
		return nil
	}

	_, err := pool.Exec(ctx,
		`INSERT INTO publications (title, journal_name, published_date, status, authors, description, tags, publication_url, order_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		"Analisis Pengujian Pengembangan Sistem Informasi Geografis Daerah Rawan Rob Di Kota Pekalongan [Testing Analysis of Geographic Information System Development for Rob Flood-Prone Areas in Pekalongan City]",
		"Jurnal Sains Komputer & Informatika (J-SAKTI)",
		"Mar 2023",
		"Published",
		"Mohammad Rasas Aufar, Pradana Ananda Raharja",
		"This study applied the Agile method to design a Geographic Information System (GIS) for mapping tidal flood-prone (rob) areas along the northern coastal zone of Pekalongan City. The GIS integration was structured to improve identification of regional vulnerability levels and to present tidal threat data to the public and relevant stakeholders in a visual and accurate format.",
		mustJSON([]string{"Geographic Information System", "Agile Method", "Tidal Flooding", "Regional Mapping"}),
		"https://garuda.kemdiktisaintek.go.id/documents/detail/3418708",
		0,
	)
	if err != nil {
		return err
	}
	log.Println("  ✅ publications seeded")
	return nil
}

func seedContactInfo(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "contact_info") {
		log.Println("  ⏭️  contact_info already seeded, skipping")
		return nil
	}

	_, err := pool.Exec(ctx,
		`INSERT INTO contact_info (email, phone, whatsapp_url, github_url, linkedin_url, instagram_url, location, contact_description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		"rasasaufar4@gmail.com",
		"(+62) 85326775595",
		"https://wa.me/085326775595",
		"https://github.com/rasasaufar",
		"https://www.linkedin.com/in/rasasaufar/",
		"https://instagram.com/rasasaufar",
		"Indonesia",
		"Feel free to reach out for collaboration, opportunities, or just to say hello.",
	)
	if err != nil {
		return err
	}
	log.Println("  ✅ contact_info seeded")
	return nil
}

func seedSiteSettings(ctx context.Context, pool *pgxpool.Pool) error {
	if tableHasData(ctx, pool, "site_settings") {
		log.Println("  ⏭️  site_settings already seeded, skipping")
		return nil
	}

	_, err := pool.Exec(ctx,
		`INSERT INTO site_settings (site_title, meta_description, theme_mode)
		VALUES ($1, $2, $3)`,
		"Rasas Aufar - IT Implementation Portfolio",
		"IT Implementation Professional experienced in government technology delivery, infrastructure operations, and applied data analysis.",
		"dark",
	)
	if err != nil {
		return err
	}
	log.Println("  ✅ site_settings seeded")
	return nil
}
