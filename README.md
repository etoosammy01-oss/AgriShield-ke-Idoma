# 🌾 Agro-Shield

> **Protecting Every Harvest. Connecting Every Farmer.**

Agro-Shield  is a digital agriculture platform designed to reduce post-harvest losses, improve market access, and empower farmers across Idoma communities and beyond.

Built for the **Idoma Centenary Plus Hackathon 2026**, the platform addresses one of Benue State's biggest agricultural challenges: helping farmers store, manage, and sell their produce at fair market prices while providing access to intelligent farming assistance.

---

# 📌 Problem Statement

Benue State is one of Nigeria's largest producers of agricultural products, yet many farmers continue to experience:

- High post-harvest losses
- Poor access to verified buyers
- Unstable market prices
- Limited storage management
- Lack of modern digital farming tools
- Poor access to agricultural information

Many smallholder farmers are forced to sell immediately after harvest at very low prices because they lack information and access to larger markets.

Agro-Shield aims to bridge this gap.

---

# 💡 Our Solution

Agro-Shield provides farmers with one platform where they can:

- Register as farmers or buyers
- Manage stored produce
- Connect directly with buyers
- Access an online agricultural marketplace
- Receive AI-powered farming assistance
- Track farming activities through a personalized dashboard

---

# 🚀 Current Features

## ✅ Landing Page

- Modern responsive homepage
- Project introduction
- Easy navigation
- Farmer-focused branding

---

## ✅ User Registration

Farmers can register by providing:

- Full Name
- Phone Number
- Password
- Location

Backend validation includes:

- Required fields
- Duplicate phone number detection
- Password hashing using bcrypt
- SQLite database storage

---

## ✅ Login Page

Secure login interface with:

- Phone number authentication
- Password verification
- Responsive design

---

## ✅ Farmer Dashboard

The dashboard currently includes:

- Welcome screen
- Product overview
- Marketplace overview
- Revenue overview
- AI Diagnoses
- Navigation system
- Quick action cards including (storang, market, ai assistant...)


#### 🤖 AI Crop Disease Assistant

Farmers will be able to:

- Upload crop images
- Detect diseases
- Receive treatment recommendations
- Learn preventive measures
---

## ✅ Buyer Dashboard

Buyers have a dedicated dashboard for:

- Marketplace access
- Purchase history
- Profile management

---

## ✅ Backend Architecture

The project follows a layered architecture.

```
Handlers
    ↓
Services
    ↓
Repositories
    ↓
SQLite Database
```

Current backend components include:

- Routing
- Middleware
- Repository Pattern
- Service Layer
- SQLite Database
- HTML Templates
- Static Asset Serving

---

# 🛠 Technology Stack

### Backend

- Go (Golang)
- net/http
- SQLite
- bcrypt
- HTML Templates

### Frontend

- HTML5
- CSS3
- JavaScript

### Version Control

- Git
- GitHub

---

# 📂 Project Structure

```
backend/
│
├── handlers/
├── internal/
│   ├── database/
│   ├── dto/
│   ├── models/
│   ├── repository/
│   ├── services/
│
├── middleware/
├── migrations/
├── render/
├── routes/
│
└── main.go
```

---

# 🌍 Target Users

- Smallholder Farmers
- Commercial Farmers
- Agricultural Cooperatives
- Product Buyers
- Agricultural Extension Workers

---

# 🎯 Hackathon Objectives

Our solution focuses on:

- Reducing post-harvest losses
- Improving market accessibility
- Increasing farmer income
- Supporting digital agriculture
- Creating a scalable platform for rural communities

---

# 🔮 Future Roadmap

The current version demonstrates the platform foundation.

Future releases will introduce the following features.

---

## 🌾 Smart Marketplace

- Direct farmer-to-buyer trading
- Verified buyer accounts
- Live product listings
- Secure transaction tracking

---

## 📈 AI Market Price Advisor

An intelligent recommendation system that helps farmers decide:

- When to sell
- Where to sell
- Expected market trends
- Price forecasting

Example:

```
Today's Yam Prices

Otukpo
₦52,000 / ton

Makurdi
₦58,000 / ton

Recommendation

Wait 2 days before selling.
Expected increase: +8%
```
---

## 📦 Smart Produce Storage

Future versions will allow farmers to:

- Track stored produce
- Monitor storage duration
- Receive spoilage alerts
- Monitor warehouse conditions

---

## 📱 USSD Support

To improve accessibility for rural communities:

- Register without internet
- Check market prices
- Receive farming tips
- Access emergency alerts

---

## 🌐 Multi-language Support

Planned support includes:

- English
- Idoma
- Egede
- Apa and more...

---

## 👨‍🌾 Cooperative Management

Farmers will be able to:

- Form cooperatives
- Aggregate produce
- Negotiate bulk sales
- Share transportation

---

## 🚚 Logistics Integration

Future releases will include:

- Transport booking
- Produce delivery tracking
- Warehouse locations
- Buyer pickup scheduling

---

## 📊 Farmer Analytics

Dashboard insights including:

- Sales history
- Revenue trends
- Produce statistics
- Storage reports

---

# 🏆 Innovation

Agro-Shield combines:

- Digital Marketplace
- Smart Storage Management
- AI Farming Assistant
- Future Price Prediction
- Farmer-Centered Design

to create an integrated agricultural ecosystem.

---

# 🌱 Sustainability

The platform is designed to:

- Increase farmer income
- Reduce food waste
- Improve food security
- Encourage digital inclusion
- Strengthen rural economies

---

# 👥 Team

Developed by Team **Agro-Shield**

Built for the **Idoma Centenary Plus Hackathon 2026**

Together, we believe technology can transform agriculture and empower every farmer.

---

# 📖 Vision

> **To become the leading digital agriculture platform connecting every farmer to better opportunities while reducing post-harvest losses through technology and innovation.**

---

# ❤️ Motto

**Protect Every Harvest. Empower Every Farmer.**