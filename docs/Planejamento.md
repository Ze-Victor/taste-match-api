# 📅 Planejamento de Atividades

## 🏁 Sprint 1 — Fundamentos do Perfil e Compatibilidade (Semana 1)

### Objetivo:
Criar a base de cadastro de perfil e início da lógica de compatibilidade.

#### ✅ Atividades:
- [ ] Criar tela de cadastro/edição de perfil (`UX/UI`)
    - Como um novo usuário, eu quero criar e editar meu perfil, para que eu possa personalizar minhas informações e começar a usar o aplicativo.

- [ ] Criar campos de preferências de locais (ex: bar, restaurante, cafeteria)
    - Como um usuário, eu quero selecionar meus tipos de locais preferidos (ex: bar, restaurante, cafeteria), para que o aplicativo me conecte com pessoas que buscam experiências semelhantes.

- [ ] Criar campos de preferências de gastronomia (ex: japonesa, italiana)
    - Como um usuário, eu quero escolher minhas preferências gastronômicas (ex: japonesa, italiana), para que eu possa encontrar pessoas com gostos culinários parecidos.

- [ ] Implementar persistência das preferências no backend
    - Como um usuário, eu quero que minhas preferências de locais e gastronomia sejam salvas, para que eu não precise inseri-las novamente a cada uso e elas sejam consideradas na busca por compatibilidade.

- [ ] Validação de campos obrigatórios
    - Como um usuário, eu quero ser avisado sobre campos obrigatórios não preenchidos durante o cadastro/edição de perfil, para que eu possa corrigir as informações e garantir que meu perfil esteja completo e válido.
    
- [ ] Desenvolver lógica inicial de compatibilidade entre perfis
    - Como o sistema, eu quero aplicar uma lógica inicial de compatibilidade com base nas preferências de locais e gastronomia dos usuários, para que eu possa identificar perfis que combinam e otimizar as sugestões.

- [ ] Criar endpoint para buscar perfis compatíveis
    - Como o sistema, eu quero ter um endpoint que retorne perfis compatíveis, para que o aplicativo possa exibir sugestões relevantes para o usuário.

- [ ] Criar layout de exibição de perfis compatíveis
    - Como um usuário, eu quero ver um layout claro e intuitivo com os perfis compatíveis sugeridos, para que eu possa facilmente visualizar e interagir com as pessoas que me interessam.

---

## 🚀 Sprint 2 — Match e Sugestão Inteligente (Semana 2)

### Objetivo:
Implementar lógica de match com intenção e sugerir locais baseados em afinidade.

#### ✅ Atividades:
- [ ] Criar botão de intenção de sair 
- [ ] Implementar regra: match só ocorre com afinidade + intenção dos dois
- [ ] Registrar matches no banco de dados
- [ ] Criar notificação de match para ambos usuários
- [ ] Criar endpoint para cruzamento de gostos entre perfis
- [ ] Criar lógica de sugestão automática de local com base em gostos comuns
- [ ] Exibir sugestão de local na tela de match

---

## 🎯 Sprint 3 — Agendamento do Encontro (Semana 3)

### Objetivo:
Permitir aos usuários combinar dia e hora e confirmar encontro.

#### ✅ Atividades:
- [ ] Criar seletor de horários disponíveis no perfil
- [ ] Sugerir horários com base no tipo de local (bar, café etc.)
- [ ] Permitir sugestão e escolha entre múltiplos horários
- [ ] Criar tela de confirmação do local e horário
- [ ] Implementar confirmação mútua (ambos confirmam o encontro)
- [ ] Registrar encontro final no sistema
- [ ] Notificar os usuários sobre o encontro agendado
- [ ] (Opcional) Gerar evento no calendário (integração futura)

---
