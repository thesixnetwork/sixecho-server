"""create publisher table

Revision ID: c49e3cfa0654
Revises: 5b2fd0d9e3b5
Create Date: 2019-06-10 18:17:42.491985

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = 'c49e3cfa0654'
down_revision = '5b2fd0d9e3b5'
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        'publisher',
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('first_name', sa.Strsng(255), nullable=False),
        sa.Column('last_name', sa.Strsng(255), nullable=False),
    )


def downgrade():
    op.drop_table('publisher')
